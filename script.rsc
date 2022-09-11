{
    /ip firewall address-list remove [/ip firewall address-list find list="ru"]
    /ipv6 firewall address-list remove [/ipv6 firewall address-list find list="ru"]
    :local ipv4 [/tool fetch url=https://raw.githubusercontent.com/dyrkin/block_ru_subnets/main/ipv4 as-value output=user];
    :if ($ipv4->"status" = "finished") do={
        :local index ($ipv4->"data");
        :local indexLen [:len $index];
        :local lineEnd 0;
        :local parititonName "";
        :local lastEnd 0;

        :while ($lineEnd < $indexLen) do={
            :set lineEnd [:find $index "\n" $lastEnd];
            :if ([:len $lineEnd] = 0) do={
                :set lineEnd $indexLen;
            }
            :set parititonName [:pick $index $lastEnd $lineEnd];
            :set lastEnd ($lineEnd + 1);
            :local parititonNameLen [:len $parititonName];
            :if ($parititonNameLen != 0) do={
                :local ipv4 [/tool fetch url="https://raw.githubusercontent.com/dyrkin/block_ru_subnets/main/$parititonName" as-value output=user];
                :if ($ipv4->"status" = "finished") do={
                    :local partition ($ipv4->"data");
                    :local partitionLen [:len $partition];
                    :local lineEndInner 0;
                    :local cidr "";
                    :local lastEndInner 0;

                    :while ($lineEndInner < $partitionLen) do={
                        :set lineEndInner [:find $partition "\n" $lastEndInner];
                        :if ([:len $lineEndInner] = 0) do={
                            :set lineEndInner $partitionLen;
                        }
                        :set cidr [:pick $partition $lastEndInner $lineEndInner];
                        :set lastEndInner ($lineEndInner + 1);
                        :local cidrLen [:len $cidr];
                        :if ($cidrLen != 0) do={
                            /ip firewall address-list add list=ru address=$cidr
                        }
                    }
                }
            }
        }
    }
     :local ipv6 [/tool fetch url=https://raw.githubusercontent.com/dyrkin/block_ru_subnets/main/ipv6 as-value output=user];
    :if ($ipv6->"status" = "finished") do={
        :local index ($ipv6->"data");
        :local indexLen [:len $index];
        :local lineEnd 0;
        :local parititonName "";
        :local lastEnd 0;

        :while ($lineEnd < $indexLen) do={
            :set lineEnd [:find $index "\n" $lastEnd];
            :if ([:len $lineEnd] = 0) do={
                :set lineEnd $indexLen;
            }
            :set parititonName [:pick $index $lastEnd $lineEnd];
            :set lastEnd ($lineEnd + 1);
            :local parititonNameLen [:len $parititonName];
            :if ($parititonNameLen != 0) do={
                :local ipv6 [/tool fetch url="https://raw.githubusercontent.com/dyrkin/block_ru_subnets/main/$parititonName" as-value output=user];
                :if ($ipv6->"status" = "finished") do={
                    :local partition ($ipv6->"data");
                    :local partitionLen [:len $partition];
                    :local lineEndInner 0;
                    :local cidr "";
                    :local lastEndInner 0;

                    :while ($lineEndInner < $partitionLen) do={
                        :set lineEndInner [:find $partition "\n" $lastEndInner];
                        :if ([:len $lineEndInner] = 0) do={
                            :set lineEndInner $partitionLen;
                        }
                        :set cidr [:pick $partition $lastEndInner $lineEndInner];
                        :set lastEndInner ($lineEndInner + 1);
                        :local cidrLen [:len $cidr];
                        :if ($cidrLen != 0) do={
                            /ipv6 firewall address-list add list=ru address=$cidr
                        }
                    }
                }
            }
        }
    }

    /ip firewall filter remove [find comment="ru_subnets"]
    /ipv6 firewall filter remove [find comment="ru_subnets"]
    /ip firewall filter add action=drop chain=input src-address-list=ru comment="ru_subnets"
    /ip firewall filter add action=drop chain=forward dst-address-list=ru comment="ru_subnets"
    /ipv6 firewall filter add action=drop chain=input src-address-list=ru comment="ru_subnets"
    /ipv6 firewall filter add action=drop chain=forward dst-address-list=ru comment="ru_subnets"
}

{
    /ip firewall address-list remove [/ip firewall address-list find list="ru"]
    /ipv6 firewall address-list remove [/ipv6 firewall address-list find list="ru"]
    :local ipv4 [/tool fetch url=https://raw.githubusercontent.com/dyrkin/ru_subnets/main/ipv4 as-value output=user];
    :if ($ipv4->"status" = "finished") do={
        :local content ($ipv4->"data");
        :local contentLen [:len $content];
        :local lineEnd 0;
        :local line "";
        :local lastEnd 0;

        :while ($lineEnd < $contentLen) do={
            :set lineEnd [:find $content "\n" $lastEnd];
            :if ([:len $lineEnd] = 0) do={
                :set lineEnd $contentLen;
            }
            :set line [:pick $content $lastEnd $lineEnd];
            :set lastEnd ($lineEnd + 1);
            :local lineLen [:len $line];
            :if ($lineLen != 0) do={
                /ip firewall address-list add list=ru address=$line
            }
        }
    }
    :local ipv6 [/tool fetch url=https://raw.githubusercontent.com/dyrkin/ru_subnets/main/ipv6 as-value output=user];
    :if ($ipv6->"status" = "finished") do={
        :local content ($ipv6->"data");
        :local contentLen [:len $content];
        :local lineEnd 0;
        :local line "";
        :local lastEnd 0;

        :while ($lineEnd < $contentLen) do={
            :set lineEnd [:find $content "\n" $lastEnd];
            :if ([:len $lineEnd] = 0) do={
                :set lineEnd $contentLen;
            }
            :set line [:pick $content $lastEnd $lineEnd];
            :set lastEnd ($lineEnd + 1);
            :local lineLen [:len $line];
            :if ($lineLen != 0) do={
                /ipv6 firewall address-list add list=ru address=$line
            }
        }
    }
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -9223372036854775808
    var v3 uint8 = 137
    switch 127 {
    case 378:
        v2 = v2
    }
    switch 127 {
    case -9223372036854775808:
        if true {
            var v4 bool = (true && true)
            _ = v4
        } else {
            switch (((("hello" + "foo bar") + "a") + "") + "\n\t") {
            case "foo bar":
                var v5 uint8 = (v3 % ((v3) | 1))
                _ = v5
            case "":
                var v6 uint8 = 127
                _ = v6
            }
            for fi7 := 0; fi7 < 18; fi7++ {
                v2 += 0
                v3 += 76
                _ = fi7
            }
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 float64 = float64(1.5)
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch (^(fi8)) {
                    case 309:
                        v3 = v3
                    case 9223372036854775807:
                        v2 = (v2 >> ((58) & 63))
                    case 2147483647:
                        switch ("\n\t" + ("hello" + ("\n\t" + ("foo bar" + "")))) {
                        case "\n\t":
                            v11 := GenS1{F0: fi10, F1: -414}
                            _ = v11
                        }
                    }
                    v12 := GenS2{F0: 934}
                    v2 -= (v2 - v2)
                    _ = fi10
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
        }
    }
    fmt.Println(v3)
    v3 += 128
    if (!("hello" != "a")) {
        v13 := []int{32768, -9223372036854775808}
        _ = v13
    }
    var v14 uint16 = 2
    if ("\n\t" < "") {
        var v15 int32 = 1
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            switch "hello" {
            case "\n\t":
                var v17 int = (fi16 << ((65535) & 63))
                _ = v17
            }
        } else {
            v1[("hello" + "foo bar")] = ("hello" + (((("hello" + "a") + "a") + "\n\t") + ""))
            switch 2147483647 {
            case -2147483648:
                for fi18 := 0; fi18 < 18; fi18++ {
                    v3 = v3
                    fmt.Println(v2)
                    _ = fi18
                }
            }
        }
        var v19 int32 = -1
        var v20 uint16 = v14
        _ = fi16
        _ = v19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}

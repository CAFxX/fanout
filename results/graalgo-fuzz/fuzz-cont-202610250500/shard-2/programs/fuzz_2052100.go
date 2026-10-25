package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int8 = 106
    var v2 uint16 = 916
    var v3 uint64 = 2
    var v4 uint = 719
    var v5 int16 = 1
    var v6 uint64 = v3
    if false {
        v7 := []int{0, -2147483648, 720, 337}
        _ = v7
    }
    v8 := GenS1{F0: 244}
    switch "foo bar" {
    case "a":
        for fi9 := 0; fi9 < 4; fi9++ {
            v3 -= v3
            v5 += 1
            v5 = v5
            _ = fi9
        }
    case "\n\t":
        v10 := []int{}
        _ = v10
    case "foo bar":
        v2 = 899
    }
    var v11 float32 = 0.0
    fmt.Println(v4, v4)
    var v12 uint64 = (v3 / (((^(v3))) | 1))
    if (-9223372036854775808 >= 42) {
        switch ((("a" + ("\n\t" + ("" + "hello"))) + "") + ("" + ("hello" + ""))) {
        case "":
            for fi13 := 0; fi13 < 14; fi13++ {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    fmt.Println(v12, v1, v6)
                case 32768:
                    v14 := GenS1{F0: (v4 * 65535)}
                    _ = v14
                }
                if true {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v16 := []int{532}
                        _ = fi15
                        _ = v16
                    }
                    v17 := []int{493}
                    _ = v17
                }
                var v18 uint32 = 4294967295
                _ = fi13
                _ = v18
            }
        case "foo bar":
            var v19 uint = (^(v4))
            _ = v19
        case "\n\t":
            v4 += v4
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            v3 -= v12
            if (!false) {
                fmt.Println(v5, v1)
                switch ("foo bar" + ("\n\t" + "foo bar")) {
                case "":
                    fmt.Println(v12)
                case "a":
                    fmt.Println(v11, v5, v3)
                }
            } else {
                switch "foo bar" {
                case "":
                    if true {
                        v21 := []int{(fi20 & 1), (fi20 / ((fi20) | 1))}
                        var v22 bool = ((false && true) && true)
                        if (v22 || true) {
                            if v22 {
                                var v23 uint32 = 2
                                _ = v23
                            }
                            switch (fi20 ^ (fi20 % ((827) | 1))) {
                            case 2147483647:
                                v11 = 1e10
                            case -2147483648:
                                v24 := GenS1{F0: 2}
                                _ = v24
                            }
                        } else {
                            fmt.Println(v22, v4)
                        }
                        _ = v21
                        _ = v22
                    } else {
                        v25 := []int{(fi20 >> ((v4) & 63)), 32768, 2147483647, (fi20 - fi20)}
                        _ = v25
                    }
                }
                var v26 float64 = float64(v11)
                _ = v26
            }
            _ = fi20
        }
    }
    v3 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

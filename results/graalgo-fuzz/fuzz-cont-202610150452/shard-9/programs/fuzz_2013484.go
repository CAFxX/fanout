package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }

func main() {
    var v1 bool = (!(31 < 255))
    v2 := []int{-9223372036854775808}
    var v3 float32 = 0.0
    var v4 int16 = -1
    v5 := []string{"\n\t"}
    v6 := []string{"a"}
    v7 := GenS1{F0: 0, F1: 946, F2: -9223372036854775808}
    v1 = v1
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v1 = false
            fmt.Println(v4, fi9)
            v4 -= v4
            _ = fi9
        }
        var v10 bool = (!(-128 <= -1))
        _ = fi8
        _ = v10
    }
    var v11 bool = false
    v3 -= v3
    v4 = 32767
    if true {
        switch "\n\t" {
        case "\n\t":
            var v12 int8 = 87
            _ = v12
        case "a":
            var v13 int16 = v4
            _ = v13
        default:
            v14 := GenS1{F0: -77, F1: 0, F2: 1}
            _ = v14
        }
        if (32767 < 859) {
            switch -9223372036854775808 {
            case 32767:
                fmt.Println(v1)
            case 9223372036854775807:
                v11 = (-2147483648 == 216)
            case -1:
                v15 := GenS1{F0: 1, F1: 851, F2: -9223372036854775808}
                _ = v15
            }
            if len(v6) > 0 {
                v6[0] = ("foo bar" + ("\n\t" + ""))
            }
            v3 = ((v3 - v3) + v3)
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                v4 = v4
                var v17 float32 = 0.1
                for fi18 := 0; fi18 < 1; fi18++ {
                    switch (fi18 % ((fi16) | 1)) {
                    case 314:
                        _ = len(v5)
                    }
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                if v11 {
                    fmt.Println(v4, v11, v1)
                    var v20 float32 = (1.5 - float32((1.5 + 1e10)))
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                        _ = fi21
                    }
                }
                _ = fi19
            }
            var v22 string = "hello"
            _ = v22
        }
    }
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    v4 = v4
    v5 = append(v5, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

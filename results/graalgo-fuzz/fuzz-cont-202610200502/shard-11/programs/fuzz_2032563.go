package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 int64, p1 int16, p2 string) uint {
    v6 := GenS1{F0: 254, F1: p1, F2: 42}
    var v7 uint64 = 42
    p0 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 561
}
func genFn1() int16 {
    var v8 bool = false
    switch 1 {
    case 32768:
        var v9 int = -930
        _ = v9
    }
    v10 := []int{-2147483648}
    _ = v10
    _ = v8
    return -32768
}

func main() {
    var v1 uint = 888
    var v2 int = 128
    var v3 float32 = -2.25
    v4 := make(map[string]bool)
    v5 := []string{("" + ((("" + "foo bar") + "a") + ("a" + "a"))), "hello", "foo bar"}
    fmt.Println(v1)
    v5 = append(v5, (("foo bar" + ("" + "\n\t")) + "foo bar"))
    var v11 int32 = 2147483647
    _ = len(v5)
    _ = len(v4)
    v11 = v11
    var v12 string = "foo bar"
    var v13 int64 = -470
    switch "\n\t" {
    case "a":
        v13 -= (v13 | (^(v13)))
    case "\n\t":
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int32 = (v11 >> ((713) & 63))
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v2)
    if (true || (!(96 > 229))) {
        switch v2 {
        case -28:
            if (true && false) {
                if (-432 < (v2 | (v2 ^ 1))) {
                    v2 = v2
                    v2 = (v2 ^ 128)
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(v2, v1)
                    _ = fi16
                }
            } else {
                for fi17 := 0; fi17 < 7; fi17++ {
                    fmt.Println(v1, v2, fi17)
                    var v18 bool = (!(false && ((fi17 ^ -9223372036854775808) == 2147483647)))
                    _ = fi17
                    _ = v18
                }
            }
        default:
            switch ("" + v12) {
            case "foo bar":
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := GenS1{F0: 128, F1: int16(v11), F2: 100}
                    if (!((v3 >= v3) || (v13 >= -614))) {
                        v3 += float32(0.0)
                    }
                    var v21 uint16 = 65535
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                v12 = ""
            }
        }
        _ = len(v4)
    }
    v11 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

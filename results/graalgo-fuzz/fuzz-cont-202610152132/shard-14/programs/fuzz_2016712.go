package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]string)
    v3 := []uint32{1}
    v1 = v1
    var v4 uint8 = 1
    v3 = append(v3, 204)
    switch 0 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: -1}
            for fi7 := 0; fi7 < 8; fi7++ {
                fmt.Println(v4, v1, fi5)
                _ = fi7
            }
            if false {
                v4 -= (v4 * v4)
                if len(v3) > 0 {
                    v3[0] = 622
                }
            } else {
                v1 = v1
                for fi8 := 0; fi8 < 20; fi8++ {
                    fmt.Println(fi8)
                    var v9 uint8 = uint8(v1)
                    var v10 int32 = -1
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            _ = fi5
            _ = v6
        }
    case 128:
        var v11 bool = false
        _ = v11
    default:
        v4 -= v4
    }
    var v12 int = -1
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS2{F0: (v1 << ((v1) & 63))}
        v4 += 128
        v15 := GenS2{F0: v1}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: "foo bar"}
    switch "\n\t" {
    case "\n\t":
        if true {
            for fi17 := 0; fi17 < 12; fi17++ {
                if ((v4 % (((v4 & (^(v4)))) | 1)) <= 0) {
                    var v18 uint64 = 1
                    _ = v18
                } else {
                    v19 := v2["foo bar"]
                    _ = v19
                }
                v1 = 786
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi17
            }
        }
    case "a":
        v2["foo bar"] = ("\n\t" + ((("\n\t" + "a") + (("" + "") + "hello")) + "foo bar"))
    default:
        for fi20 := 0; fi20 < 15; fi20++ {
            switch (-(v12)) {
            case -786:
                var v21 int8 = 4
                _ = v21
            case -9223372036854775808:
                var v22 string = (("" + ((("hello" + ("\n\t" + "hello")) + "foo bar") + ("hello" + "foo bar"))) + (("hello" + "\n\t") + "a"))
                _ = v22
            case 379:
                v23 := GenS3{F0: "\n\t"}
                _ = v23
            }
            v4 = v4
            _ = fi20
        }
    }
    var v24 int16 = -248
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}

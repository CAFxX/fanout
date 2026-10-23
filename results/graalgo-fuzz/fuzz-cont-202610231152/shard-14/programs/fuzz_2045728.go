package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 uint64) int16 {
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 12; fi5++ {
        switch "\n\t" {
        case "a":
            v6 := []int{fi5, (fi5 + fi5)}
            _ = v6
        default:
            if ((p0 & (^(p0))) == -76) {
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 uint16 = 255
                    fmt.Println(fi7)
                    v9 := GenS1{F0: v4}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                p1 = (p1 - (p1 << ((65535) & 63)))
                switch (fi5 + fi5) {
                case 128:
                    var v10 bool = false
                    _ = v10
                default:
                    v4 = ""
                }
            } else {
                var v11 float32 = (3.14159 * (1e10 / 1.0))
                _ = v11
            }
        }
        var v12 int32 = 42
        if ((p0 * 0) != 125) {
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v12)
                fmt.Println(p1, p0, p0)
                _ = fi13
            }
            switch -9223372036854775808 {
            case 32767:
                v12 += (v12 >> ((62) & 63))
            case -134:
                v12 -= (v12 * 42)
            default:
                v14 := GenS1{F0: "foo bar"}
                _ = v14
            }
        } else {
            var v15 uint16 = 255
            v15 = v15
            _ = v15
        }
        _ = fi5
        _ = v12
    }
    v16 := []int{9223372036854775807, 127, 127, -2147483648}
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    return 243
}
func genFn1() uint32 {
    var v17 uint64 = 568
    var v18 uint64 = v17
    _ = v17
    _ = v18
    return 656
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{}
    var v3 string = (("hello" + (("hello" + "foo bar") + "a")) + (("hello" + "hello") + "\n\t"))
    v19 := GenS2{F0: 1}
    v3 = (v3 + "a")
    var v20 uint = 347
    var v21 float64 = (0.0 / (1.5 + 1e10))
    if (-1 > -398) {
        if false {
            for fi22 := 0; fi22 < 20; fi22++ {
                v20 = (v20 ^ 324)
                v23 := GenS3{F0: v20}
                _ = fi22
                _ = v23
            }
        } else {
            var v24 int = 478
            _ = v24
        }
        if (478 < 65535) {
            var v25 int64 = 9223372036854775807
            if true {
                var v26 int32 = -2147483648
                _ = v26
            }
            _ = v25
        }
    }
    var v27 int8 = 1
    v27 = (^(v27))
    v21 = (1.5 + float64(v1))
    var v28 int32 = 2147483647
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    v29 := GenS4{F0: 4294967295, F1: int(v20), F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}

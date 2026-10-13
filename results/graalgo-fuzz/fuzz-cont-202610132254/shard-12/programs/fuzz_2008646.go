package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 42
    var v3 float64 = 3.14159
    if (810 == 65535) {
        v2 = 1
        v4 := v1[0]
        _ = v4
    } else {
        var v5 string = "\n\t"
        _ = v5
    }
    v6 := GenS1{F0: 9223372036854775807}
    var v7 uint32 = 2147483648
    var v8 int = -2147483648
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: (fi10 + -817)}
        switch (fi10 | -2147483648) {
        case -674:
            fmt.Println(fi10)
        case -525:
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(fi10, fi12)
                fmt.Println(v8, v9)
                _ = fi12
            }
        }
        _ = fi10
        _ = v11
    }
    if false {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch fi13 {
            case 0:
                v14 := GenS2{F0: -410, F1: v9}
                _ = v14
            case 32768:
                if ((100 >= 255) && ((58 < (v2 << ((4294967295) & 63))) && true)) {
                    v9 -= v9
                    var v15 int32 = -1
                    _ = v15
                } else {
                    v7 -= (v7 ^ uint32(v9))
                    var v16 uint = 65535
                    var v17 uint64 = 942
                    _ = v16
                    _ = v17
                }
            case 128:
                if (false && ((42 == 161) || false)) {
                    var v18 int64 = -1
                    _ = v18
                }
            }
            v7 = (v7 << ((252) & 63))
            _ = fi13
        }
        if (0 >= 1) {
            v1[2147483647] = (((("foo bar" + "") + ("foo bar" + ("" + ((("" + (("hello" + "a") + "hello")) + (("a" + "hello") + "\n\t")) + ("foo bar" + "\n\t"))))) + "hello") + ("hello" + (("hello" + "") + (("\n\t" + "") + "hello"))))
            var v19 uint16 = v9
            _ = v19
        }
        if true {
            if true {
                v20 := GenS1{F0: 127}
                _ = v20
            } else {
                switch "a" {
                case "\n\t":
                    v21 := GenS1{F0: (v8 / (((v8 % (((v8 % (((^(v8))) | 1))) | 1))) | 1))}
                    _ = v21
                case "a":
                    v8 = v8
                }
                v7 -= v7
            }
        } else {
            if ((2 <= 487) || ((((float32(((v3 * (((v3 - (0.0 - v3)) - v3) * 0.0)) / v3)) * 0.1) * 1.5) - 0.0) <= float32(v3))) {
                if (true && (-97 == 42)) {
                    var v22 uint8 = 137
                    var v23 int = (v8 | 4294967295)
                    var v24 int = v8
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    v8 = -1
                    v8 = (v8 << ((387) & 63))
                }
                fmt.Println(v8, v8, v7)
            }
            var v25 float32 = (1.5 / (((0.1 * (0.1 + 0.0)) - ((0.0 - 0.1) / 3.14159)) + 0.1))
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS2{F0: 42, F1: (v9 - (v9 / ((v9) | 1)))}
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 string = ("\n\t" + "a")
                    _ = fi28
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    }
    fmt.Println(v9, v9)
    for fi30 := 0; fi30 < 4; fi30++ {
        switch 128 {
        case 128:
            fmt.Println(v8, v8)
        case 4294967295:
            fmt.Println(fi30, v8)
        case -2147483648:
            v31 := GenS1{F0: 365}
            _ = v31
        default:
            switch (fi30 * (fi30 * 32767)) {
            case -1:
                if ((("" + ((("hello" + ((("" + "hello") + "a") + "hello")) + ("a" + ("foo bar" + ("foo bar" + "a")))) + "hello")) + "\n\t") > "a") {
                    var v32 string = "\n\t"
                    for fi33 := 0; fi33 < 20; fi33++ {
                        fmt.Println(v2)
                        if ((v2 % ((v2) | 1)) < 0) {
                            var v34 int16 = int16(v9)
                            _ = v34
                        }
                        _ = fi33
                    }
                    v35 := []int{(fi30 ^ (v8 + 32767)), (^(v8))}
                    _ = v32
                    _ = v35
                }
            case 2147483647:
                if true {
                    var v36 bool = ((-128 > 100) && true)
                    _ = v36
                } else {
                    v37 := []int{(v8 & -2147483648), (v8 >> ((0) & 63)), fi30}
                    _ = v37
                }
            case 42:
                if true {
                    var v38 uint8 = 0
                    v39 := []int{fi30, 4294967295}
                    var v40 string = "a"
                    _ = v38
                    _ = v39
                    _ = v40
                }
            default:
                v41 := GenS3{F0: (v2 - v2)}
                _ = v41
            }
        }
        var v42 float64 = v3
        var v43 uint16 = v9
        _ = fi30
        _ = v42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

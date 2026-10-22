package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int }
func genFn0() int8 {
    var v5 int = 9223372036854775807
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v5, v5)
        _ = fi6
    }
    v7 := GenS1{F0: 1000}
    v5 -= 32768
    _ = v5
    _ = v7
    return 0
}
func genFn1() uint16 {
    if (1 < -57) {
        switch (("hello" + ((("" + "a") + (("\n\t" + "a") + "hello")) + "")) + "hello") {
        case "\n\t":
            v8 := GenS1{F0: 80}
            _ = v8
        case "":
            if false {
                if false {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        fmt.Println(fi9)
                        _ = fi9
                    }
                    var v10 float32 = ((float32(1.5) + float32((-2.25 * 0.0))) * -2.25)
                    if (1 < 1) {
                        var v11 string = ("" + "\n\t")
                        v12 := []int{}
                        _ = v11
                        _ = v12
                    }
                    _ = v10
                }
                v13 := []int{141, 127}
                var v14 float32 = 1e10
                _ = v13
                _ = v14
            }
        case "a":
            for fi15 := 0; fi15 < 2; fi15++ {
                switch "" {
                case "hello":
                    var v16 float64 = (((0.1 / 1.0) * -2.25) * (-2.25 + -2.25))
                    _ = v16
                }
                var v17 int64 = 0
                _ = fi15
                _ = v17
            }
        default:
            var v18 uint8 = 127
            _ = v18
        }
        var v19 float64 = (-2.25 / (3.14159 / float64(1e10)))
        _ = v19
    } else {
        var v20 uint16 = 475
        var v21 int16 = 32767
        v22 := []int{0, -534}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: 154, F1: 227}
    var v24 int64 = 1
    var v25 float64 = 3.14159
    v25 += v25
    _ = v23
    _ = v24
    _ = v25
    return 2
}

func main() {
    var v1 string = ""
    var v2 uint64 = 42
    v3 := []int{248}
    var v4 bool = (!(v2 >= (v2 << ((255) & 63))))
    for fi26 := 0; fi26 < 14; fi26++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v27 := GenS2{F0: 42, F1: 128}
        _ = fi26
        _ = v27
    }
    fmt.Println(v1, v4)
    if (true && (128 <= uint8(v2))) {
        v28 := GenS2{F0: 2, F1: 127}
        _ = v28
    } else {
        v29 := GenS3{F0: 1000, F1: 0}
        _ = v29
    }
    v4 = ("foo bar" == v1)
    if (!true) {
        v30 := GenS1{F0: 32767}
        v1 = v1
        _ = v30
    } else {
        _ = len(v3)
        v31 := GenS4{F0: v4, F1: 42}
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch 991 {
            case 986:
                v32 := GenS4{F0: true, F1: 1}
                _ = v32
            default:
                for fi33 := 0; fi33 < 18; fi33++ {
                    var v34 int8 = 100
                    _ = fi33
                    _ = v34
                }
            }
            v1 = ("a" + "foo bar")
        }
        _ = v31
    }
    v1 = "foo bar"
    v35 := GenS2{F0: 71, F1: 210}
    if (27 <= 2) {
        if ((671 <= 2147483647) && (v4 && (1000 == -158))) {
            var v36 bool = v4
            _ = v36
        }
        if len(v3) > 0 {
            v3[0] = -168
        }
        var v37 uint8 = 253
        _ = v37
    }
    if false {
        var v38 bool = v4
        v39 := GenS4{F0: true, F1: 0}
        _ = v38
        _ = v39
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 -= (v2 + v2)
    }
    fmt.Println(v1)
    var v40 uint64 = v2
    var v41 string = v1
    v2 = 661
    fmt.Println(v2, v1, v4)
    fmt.Println(v40)
    var v42 int64 = -513
    v43 := GenS5{F0: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}

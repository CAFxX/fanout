package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint; F2 uint }
func genFn0(p0 int8, p1 int32) bool {
    fmt.Println(p0, p0)
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := []int{fi4, -9223372036854775808}
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return (false && (-1 >= (p1 >> ((4294967295) & 63))))
}

func main() {
    var v1 uint8 = 132
    var v2 int8 = 100
    var v3 uint = 42
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 string = ""
        var v8 string = ""
        if false {
            v9 := GenS1{F0: v1}
            _ = v9
        } else {
            var v10 float32 = float32((float64(0.0) - ((-2.25 * (0.1 * 0.0)) + (0.1 / ((float64(3.14159) / 0.1) + -2.25)))))
            v11 := []int{(^(fi6))}
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        var v12 int16 = -32768
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint32 = 0
            switch fi13 {
            case 128:
                if false {
                    if true {
                        var v15 int64 = -1
                        var v16 uint8 = 2
                        var v17 string = "hello"
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                } else {
                    v12 -= v12
                }
            case 1:
                var v18 int8 = (v2 >> ((655) & 63))
                _ = v18
            case 4294967295:
                var v19 uint32 = v14
                _ = v19
            default:
                var v20 uint = v3
                _ = v20
            }
            if ((1.5 + -2.25) < 0.1) {
                v2 = int8(v14)
                v3 = v3
                v14 += (v14 % ((v14) | 1))
            }
            _ = fi13
            _ = v14
        }
        var v21 float32 = (((1e10 / 0.1) + -2.25) / (float32(3.14159) + (-2.25 / ((((-2.25 + (-2.25 * 0.1)) / float32(3.14159)) * 1e10) + float32((3.14159 / ((float64((1e10 / float32(3.14159))) + 3.14159) + 0.1)))))))
        _ = v12
        _ = v21
    }
    v22 := []int{1, 480}
    v23 := GenS1{F0: (v1 & (v1 + 2))}
    if (0.1 < 1.5) {
        var v24 uint16 = 2
        fmt.Println(v3, v24)
        v2 = 1
        _ = v24
    }
    fmt.Println(v2, v1)
    if false {
        if false {
            if len(v22) > 0 {
                v22[0] = 32767
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                fmt.Println(v2, v3)
                _ = fi25
            }
            v26 := GenS1{F0: v1}
            _ = v26
        }
        v1 += v1
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v28 := GenS2{F0: 1}
        var v29 string = ""
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int = int(v2)
    v3 = (v3 & (v3 & (v3 | (v3 - 295))))
    if false {
        _ = len(v22)
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v22 = append(v22, int(v1))
    fmt.Println(v30, v1, v2)
    fmt.Println(v30)
    switch int(v2) {
    case 4294967295:
        if (true && (("" != "") || true)) {
            v2 = -128
        }
    case -2147483648:
        if len(v22) > 0 {
            v22[0] = 2147483647
        }
    case 839:
        v31 := GenS3{F0: 1, F1: v3, F2: (v3 << (((v3 / (((v3 ^ v3)) | 1))) & 63))}
        _ = v31
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        fmt.Println(v1)
        if ((false && false) && ((("hello" + "\n\t") + "a") != "\n\t")) {
            if ((!(-412 == -211)) && true) {
                for fi33 := 0; fi33 < 20; fi33++ {
                    v30 -= v30
                    _ = fi33
                }
            }
            var v34 float32 = (((1.5 - float32(((0.1 + 3.14159) / 0.1))) + ((-2.25 + (1e10 + ((float32(0.1) - (0.0 * -2.25)) + ((3.14159 - 1.5) / (1e10 * (float32(3.14159) + 0.1)))))) + ((-2.25 * 1e10) / 3.14159))) - 3.14159)
            _ = v34
        } else {
            var v35 int = int(v3)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = v35
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
}

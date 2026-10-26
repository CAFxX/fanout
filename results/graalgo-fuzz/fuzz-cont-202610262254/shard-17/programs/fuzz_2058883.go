package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint8, p1 bool) int16 {
    var v6 uint = 17
    if (42 > (p0 >> ((v6) & 63))) {
        p0 = 39
        switch 0 {
        case 1:
            p1 = (((42 <= 678) || p1) && (!p1))
        }
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        if p1 {
            for fi8 := 0; fi8 < 6; fi8++ {
                p0 -= p0
                if p1 {
                    fmt.Println(v6, fi7)
                    v9 := GenS1{F0: 594}
                    v6 = 65535
                    _ = v9
                }
                _ = fi8
            }
        }
        p1 = (-1 < 0)
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int8 = 122
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    var v12 int8 = 42
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := []int{}
        var v15 float32 = (0.1 / float32(((1e10 / float64(1e10)) + ((((1.5 - (0.0 + (((3.14159 - 0.1) / ((0.1 + float64(((((-2.25 / 1.5) - ((1e10 + float32((1.5 * 3.14159))) - (1e10 / 1.5))) / 1.5) / 1e10))) * (((float64(1e10) - 1.5) + (3.14159 * (1.5 + (((float64((0.0 * 1.5)) - 1.5) + 0.1) + 1.5)))) * -2.25))) / 3.14159))) + 1.5) / 1.5) + 1.5))))
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return -606
}

func main() {
    var v1 uint = 65535
    var v2 int16 = 1000
    v3 := make(map[int]string)
    var v4 int16 = (v2 & v2)
    v5 := make(map[uint32]uint32)
    v4 -= (v2 | 32767)
    fmt.Println(v4, v4, v1)
    fmt.Println(v2)
    v16 := []int{-2147483648}
    if true {
        if ("\n\t" <= "hello") {
            v4 += (v4 << ((v1) & 63))
        } else {
            if ((-956 > -1) && (128 == 128)) {
                var v17 float32 = -2.25
                v2 -= (v2 / (((v4 - -1)) | 1))
                _ = v17
            } else {
                v4 -= 32767
            }
            v18 := GenS2{F0: true, F1: 42, F2: 394}
            _ = v18
        }
        v19 := GenS3{F0: 4294967295}
        for fi20 := 0; fi20 < 6; fi20++ {
            v5[4294967295] = 144
            _ = fi20
        }
        _ = v19
    }
    v1 -= (v1 >> (((v1 * (v1 << (((v1 << (((^(v1))) & 63))) & 63)))) & 63))
    if true {
        if (((-(v2)) >= v2) || false) {
            var v21 string = "hello"
            v22 := GenS3{F0: 18446744073709551615}
            for fi23 := 0; fi23 < 20; fi23++ {
                v4 = -1
                var v24 int8 = 100
                _ = fi23
                _ = v24
            }
            _ = v21
            _ = v22
        }
        v1 += 914
        v1 += v1
    } else {
        switch ((("a" + "\n\t") + "a") + "\n\t") {
        case "a":
            fmt.Println(v4, v1, v1)
        default:
            v25 := GenS3{F0: 811}
            _ = v25
        }
        v26 := GenS1{F0: -684}
        _ = v26
    }
    var v27 uint64 = 18446744073709551615
    v4 = v4
    var v28 int = 4294967295
    var v29 uint32 = 2
    for fi30 := 0; fi30 < 17; fi30++ {
        v5[646] = v29
        v2 += (v2 / ((v4) | 1))
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := GenS2{F0: ((false || false) && (v1 >= v1)), F1: v27, F2: 1}
            if (255 == 865) {
                v29 = v29
            }
            var v33 int = (fi31 / ((fi30) | 1))
            var v34 int = 178
            v4 = genFn0(2, (!((!true) && (fi30 < (fi31 >> ((v1) & 63))))))
            v1 = 0
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi30
    }
    v1 = uint(v27)
    if false {
        switch "foo bar" {
        case "\n\t":
            var v35 int8 = 42
            _ = v35
        }
        v36 := v3[(-(v28))]
        var v37 string = (("\n\t" + v36) + v36)
        _ = v36
        _ = v37
    }
    if false {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            v5[(v29 | v29)] = 1
            _ = fi38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

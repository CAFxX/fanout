package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 0
    var v3 string = (("hello" + ((("hello" + ("hello" + "")) + "foo bar") + "")) + ("a" + "hello"))
    var v4 uint = 181
    v5 := make(map[int]bool)
    switch v3 {
    case "a":
        v4 = 18446744073709551615
    case "":
        var v6 int32 = -516
        _ = v6
    }
    switch v3 {
    case "\n\t":
        v7 := []int{-920, 9223372036854775807, 32767}
        _ = v7
    case "":
        fmt.Println(v3, v3, v2)
    default:
        v4 += 0
    }
    var v8 uint16 = 568
    fmt.Println(v3, v2, v3)
    switch 2147483647 {
    case 895:
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 bool = true
            v11 := []int{fi9, fi9, (^(fi9))}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int64 = (v2 | (v2 & v2))
    if false {
        var v13 bool = (!(v4 == v4))
        _ = v13
    }
    v14 := GenS1{F0: v12}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint32 = 0
        _ = fi15
        _ = v16
    }
    if true {
        if true {
            v17 := GenS1{F0: v2}
            v18 := []int{405, 0}
            var v19 float64 = (0.0 - ((((0.0 + 1e10) / 1.0) - (-2.25 - 1e10)) * ((3.14159 - float64(float32(1.5))) / float64(3.14159))))
            _ = v17
            _ = v18
            _ = v19
        } else {
            v4 += 779
        }
        fmt.Println(v4, v3, v3)
        if true {
            v20 := GenS1{F0: -1}
            _ = v20
        }
    }
    if true {
        v21 := []int{32767, 0, -9223372036854775808}
        v22 := GenS1{F0: (v2 % ((-9223372036854775808) | 1))}
        v23 := GenS2{F0: 458, F1: true, F2: 352}
        _ = v21
        _ = v22
        _ = v23
    } else {
        v24 := GenS1{F0: int64(v8)}
        switch 127 {
        case 42:
            switch v3 {
            case "":
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 uint16 = 1
                    _ = fi25
                    _ = v26
                }
            case "a":
                fmt.Println(v2)
            }
        case 1:
            if false {
                if false {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        if (((((1e10 + ((1.5 + (float64(1e10) / (-2.25 / 1e10))) * 0.0)) / (-2.25 + -2.25)) / (0.0 - (3.14159 * float64(1.5)))) == (float64(0.0) - 3.14159)) || false) {
                            if (445 <= 1000) {
                                for fi28 := 0; fi28 < 11; fi28++ {
                                    v8 = 0
                                    _ = fi28
                                }
                            }
                            v3 = v3
                            v29 := GenS3{F0: 62, F1: (fi27 << (((v4 / ((v4) | 1))) & 63))}
                            _ = v29
                        }
                        if true {
                            var v30 int16 = int16(fi27)
                            delete(v1, fi27)
                            _ = len(v5)
                            _ = v30
                        }
                        var v31 int64 = (v12 / ((v12) | 1))
                        _ = fi27
                        _ = v31
                    }
                }
                v32 := []int{}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 int = fi33
                    var v35 uint8 = 128
                    fmt.Println(v2, v35, v12)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                v12 = (v2 ^ (v2 / ((v12) | 1)))
                var v36 uint8 = 255
                _ = v36
            }
        default:
            if (v12 == (v12 - (-(v2)))) {
                if false {
                    v37 := []int{32767, -1, -1, -569}
                    _ = v37
                }
            } else {
                v8 -= (v8 % (((v8 * v8)) | 1))
                var v38 int = -955
                switch v3 {
                case "hello":
                    v3 = (v3 + (("a" + v3) + "hello"))
                case "foo bar":
                    v39 := GenS2{F0: 1, F1: ((v12 * v12) > -157), F2: 776}
                    _ = v39
                case "\n\t":
                    var v40 uint8 = 255
                    _ = v40
                }
                _ = v38
            }
        }
        fmt.Println(v4, v4)
        _ = v24
    }
    v2 += v12
    v2 += (v2 * (^(v12)))
    var v41 int64 = v2
    switch -1 {
    case -828:
        v8 = 255
    case 32768:
        fmt.Println(v41, v41)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v8
}

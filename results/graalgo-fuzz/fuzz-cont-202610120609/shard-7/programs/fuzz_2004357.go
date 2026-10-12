package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int64; F2 int64 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint16 = 32767
    v6 -= v6
    _ = v6
    return 128
}
func genFn1(p0 uint32, p1 int8, p2 uint32) uint64 {
    if ((!false) || true) {
        fmt.Println(p0, p2, p0)
        v7 := []int{-9223372036854775808, 127}
        fmt.Println(p1)
        _ = v7
    }
    if (!false) {
        v8 := []int{2147483647, 32768}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch "\n\t" {
        case "":
            _ = len(v8)
        case "\n\t":
            var v9 uint16 = 255
            _ = v9
        default:
            var v10 int16 = 1000
            _ = v10
        }
        _ = v8
    } else {
        fmt.Println(p0, p2, p2)
        for fi11 := 0; fi11 < 8; fi11++ {
            if ((p0 / (((p0 ^ p0)) | 1)) <= (^(p0))) {
                v12 := GenS1{F0: 557}
                p2 -= p0
                p1 = (p1 + (p1 * 0))
                _ = v12
            } else {
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 uint = 65535
                    p2 -= 298
                    _ = fi13
                    _ = v14
                }
                v15 := GenS2{F0: 9223372036854775807, F1: 32767, F2: 0}
                _ = v15
            }
            v16 := []int{(fi11 & fi11), 772}
            _ = fi11
            _ = v16
        }
        v17 := []int{}
        _ = v17
    }
    v18 := []int{}
    if true {
        var v19 float32 = -2.25
        v20 := GenS3{F0: 1000, F1: 691}
        v21 := GenS3{F0: -1, F1: 0}
        _ = v19
        _ = v20
        _ = v21
    } else {
        fmt.Println(p0, p2)
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 18446744073709551615
}

func main() {
    var v1 bool = false
    var v2 int32 = 0
    var v3 uint64 = 364
    v4 := make(map[string]int)
    fmt.Println(v1, v3)
    var v22 float32 = (0.1 / 1.0)
    var v23 bool = (!((v1 && (3.14159 >= -2.25)) || (v1 && (127 <= 72))))
    v24 := GenS2{F0: 286, F1: 565, F2: (v3 << ((2) & 63))}
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int16 = -852
        switch (-(fi25)) {
        case 32768:
            if v23 {
                switch (fi25 / ((fi25) | 1)) {
                case 127:
                    if (v22 <= (1.5 / v22)) {
                        switch (fi25 >> ((569) & 63)) {
                        case 607:
                            v27 := GenS3{F0: (v26 & v26), F1: genFn1(0, -128, 25)}
                            _ = v27
                        default:
                            var v28 float32 = (((v22 + 3.14159) * 0.1) / 3.14159)
                            _ = v28
                        }
                        fmt.Println(v2, v22)
                    }
                default:
                    v29 := v4[""]
                    _ = v29
                }
                var v30 uint16 = 42
                v2 = v2
                _ = v30
            }
        case -2147483648:
            v31 := []int{fi25, (fi25 >> ((0) & 63)), fi25}
            _ = v31
        default:
            switch "a" {
            case "":
                fmt.Println(v23, v23)
            case "hello":
                var v32 uint = 1
                _ = v32
            case "a":
                var v33 bool = v23
                _ = v33
            default:
                v34 := []int{fi25}
                _ = v34
            }
        }
        _ = fi25
        _ = v26
    }
    var v35 uint64 = (v3 | (v3 / ((v3) | 1)))
    var v36 uint16 = 1
    var v37 uint = 0
    for fi38 := 0; fi38 < 8; fi38++ {
        v22 += v22
        if v23 {
            for fi39 := 0; fi39 < 10; fi39++ {
                var v40 int32 = v2
                if true {
                    v35 = (v35 << (((^(v37))) & 63))
                    v2 -= 1
                    switch ("" + "hello") {
                    case "hello":
                        v41 := GenS4{F0: 351, F1: 0, F2: (^(v37))}
                        _ = v41
                    case "foo bar":
                        v42 := GenS1{F0: 42}
                        _ = v42
                    default:
                        v37 = 0
                    }
                }
                _ = fi39
                _ = v40
            }
            v22 = v22
            fmt.Println(v36)
        } else {
            fmt.Println(v22, v35, v22)
            v2 -= (-(v2))
        }
        v23 = (v2 != 42)
        _ = fi38
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        switch "a" {
        case "foo bar":
            if v1 {
                v44 := GenS3{F0: 0, F1: v3}
                v45 := GenS5{F0: 0}
                v1 = (((v1 || v23) && false) || true)
                _ = v44
                _ = v45
            } else {
                v46 := []int{}
                var v47 int = (fi43 + fi43)
                _ = v46
                _ = v47
            }
        case "a":
            v48 := GenS1{F0: 32767}
            _ = v48
        case "\n\t":
            var v49 int64 = 1
            _ = v49
        }
        var v50 int32 = (v2 >> ((uint(v35)) & 63))
        if false {
            v51 := GenS6{F0: 643, F1: -9223372036854775808, F2: 42}
            if v23 {
                fmt.Println(v3, v37, v50)
                v1 = false
                v1 = (!v23)
            }
            _ = v51
        }
        _ = fi43
        _ = v50
    }
    var v52 uint8 = 127
    v3 = uint64(v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v52
}

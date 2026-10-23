package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 string }

func main() {
    var v1 string = ("" + ("a" + "foo bar"))
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 int = 1
    v5 := []uint32{2147483647}
    var v6 int = v4
    v1 = v1
    if v2 {
        if (!v2) {
            if (!true) {
                v7 := GenS1{F0: -1, F1: v2}
                var v8 float64 = 3.14159
                switch v1 {
                case "hello":
                    v6 = (v6 >> ((65535) & 63))
                }
                _ = v7
                _ = v8
            } else {
                if true {
                    fmt.Println(v6)
                    switch v4 {
                    case -486:
                        var v9 uint64 = 18446744073709551615
                        _ = v9
                    default:
                        v4 = v4
                    }
                }
                for fi10 := 0; fi10 < 19; fi10++ {
                    v6 = (^(fi10))
                    v11 := GenS1{F0: 42, F1: (337 < 32767)}
                    _ = fi10
                    _ = v11
                }
                var v12 uint = 69
                _ = v12
            }
            v13 := v3[(v4 % (((v6 ^ v4)) | 1))]
            _ = v13
        }
        fmt.Println(v2)
    }
    var v14 int64 = 0
    for fi15 := 0; fi15 < 6; fi15++ {
        if (float32(1.5) <= (1e10 * float32(0.0))) {
            if v2 {
                v16 := GenS2{F0: 42}
                _ = v16
            } else {
                var v17 uint16 = 0
                _ = v17
            }
            fmt.Println(v2, v14)
            var v18 float64 = (float64((1e10 * float32(1.5))) / -2.25)
            _ = v18
        }
        _ = fi15
    }
    v1 = "\n\t"
    fmt.Println(v1, v2)
    v19 := GenS2{F0: 1}
    v20 := GenS3{F0: v14, F1: 0, F2: (v6 != v4)}
    v21 := GenS1{F0: 70, F1: (!true)}
    var v22 uint32 = 4294967295
    var v23 uint = 18446744073709551615
    var v24 uint16 = 65535
    if v2 {
        v25 := GenS4{F0: 2147483647, F1: "foo bar"}
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint8 = 1
            switch v4 {
            case 127:
                fmt.Println(v4, v1, v23)
            }
            v27 = v27
            _ = fi26
            _ = v27
        }
        if true {
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 int32 = 42
                var v30 int16 = 1
                if true {
                    var v31 bool = ((109 >= 255) || false)
                    var v32 float32 = ((float32((1.5 * float64(-2.25))) - 1e10) / 1e10)
                    _ = v31
                    _ = v32
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = v25
    } else {
        if len(v5) > 0 {
            v5[0] = v22
        }
        switch v1 {
        case "a":
            var v33 int8 = -57
            _ = v33
        case "\n\t":
            _ = len(v3)
        default:
            v2 = (127 > v4)
        }
    }
    switch v6 {
    case -1:
        var v34 int64 = v14
        _ = v34
    case 32768:
        for fi35 := 0; fi35 < 19; fi35++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v36 uint16 = (v24 ^ (v24 % ((1) | 1)))
            fmt.Println(v36, fi35)
            _ = fi35
            _ = v36
        }
    }
    v6 -= (-(v6))
    var v37 bool = v2
    if v37 {
        for fi38 := 0; fi38 < 1; fi38++ {
            for fi39 := 0; fi39 < 20; fi39++ {
                v14 -= 173
                v40 := GenS3{F0: (v14 - v14), F1: 59, F2: ((v14 % ((v14) | 1)) > v14)}
                _ = fi39
                _ = v40
            }
            fmt.Println(v22, v2, v37)
            v41 := GenS5{F0: v1}
            _ = fi38
            _ = v41
        }
        var v42 int = (v4 ^ v4)
        v22 += (v22 - (v22 + v22))
        _ = v42
    }
    fmt.Println(v14, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

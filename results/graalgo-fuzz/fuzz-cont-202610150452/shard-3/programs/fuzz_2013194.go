package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 uint32 }
func genFn0() int16 {
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi7)
            var v8 uint = 2
            v9 := []int{135, fi7, (^(fi7)), fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        var v10 int16 = 32767
        _ = v10
    }
    v11 := []int{217}
    _ = v11
    return -320
}

func main() {
    var v1 bool = true
    v2 := []int{32767, 128, -9223372036854775808, 127}
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := v6["\n\t"]
        v14 := GenS1{F0: v13, F1: -954}
        switch v13 {
        case "hello":
            switch ((v13 + v13) + (v13 + "hello")) {
            case "\n\t":
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS1{F0: v13, F1: 494}
                    v1 = false
                    _ = fi15
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 10; fi17++ {
                    v1 = v1
                    v18 := GenS1{F0: v13, F1: 32767}
                    var v19 uint32 = 628
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        case "a":
            var v20 int16 = -1
            _ = v20
        default:
            if v1 {
                v21 := GenS2{F0: 1, F1: fi12}
                _ = v21
            } else {
                var v22 int = (fi12 / (((^(fi12))) | 1))
                v22 += -434
                fmt.Println(v1)
                _ = v22
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    if true {
        switch 42 {
        case 116:
            if (0.0 < ((0.0 / float32(1e10)) + (3.14159 * 1.5))) {
                v1 = (-32768 != -971)
                for fi23 := 0; fi23 < 1; fi23++ {
                    if len(v2) > 0 {
                        v2[0] = fi23
                    }
                    _ = fi23
                }
                v1 = ((!true) && (403 != 1))
            }
        case 2147483647:
            v24 := GenS3{F0: 1}
            _ = v24
        case 9223372036854775807:
            v25 := GenS2{F0: 0, F1: -9223372036854775808}
            _ = v25
        }
        fmt.Println(v1)
        v1 = v1
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        v1 = v1
        _ = fi26
    }
    if v1 {
        var v27 int32 = -1
        _ = v27
    } else {
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS4{F0: 933}
            v30 := GenS4{F0: 1}
            var v31 uint16 = 579
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
        v32 := GenS1{F0: "hello", F1: 0}
        var v33 uint16 = 32767
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v1)
        if true {
            for fi35 := 0; fi35 < 2; fi35++ {
                v2 = append(v2, (fi35 / ((fi34) | 1)))
                switch (fi34 << ((656) & 63)) {
                case 127:
                    var v36 bool = v1
                    _ = v36
                case -870:
                    fmt.Println(fi34, fi35)
                case -9223372036854775808:
                    var v37 int = 128
                    _ = v37
                }
                v1 = true
                v38 := GenS5{F0: 42, F1: 2}
                _ = fi35
                _ = v38
            }
            v1 = v1
        }
        if (((("" + "") + ("hello" + "hello")) + "foo bar") == ("a" + "foo bar")) {
            delete(v6, "")
        }
        _ = fi34
    }
    var v39 bool = (9223372036854775807 >= -97)
    v39 = v1
    v39 = false
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 uint8 = 1
        v2 = append(v2, fi40)
        var v42 int64 = 336
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int = -9223372036854775808
    fmt.Println(v1)
    var v44 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}

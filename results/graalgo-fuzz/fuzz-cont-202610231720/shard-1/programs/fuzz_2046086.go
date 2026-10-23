package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 string; F2 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 int16, p2 int) int16 {
    for fi4 := 0; fi4 < 6; fi4++ {
        switch ("\n\t" + "a") {
        case "":
            var v5 string = ((("a" + "foo bar") + (("hello" + "hello") + "a")) + "")
            _ = v5
        case "\n\t":
            switch (fi4 / (((^(fi4))) | 1)) {
            case 0:
                var v6 bool = true
                _ = v6
            case 524:
                var v7 float32 = (-2.25 - (3.14159 * 0.1))
                _ = v7
            case -1:
                var v8 uint32 = 65535
                _ = v8
            default:
                v9 := GenS1{F0: 1000}
                _ = v9
            }
        }
        var v10 int = (p2 << ((p0) & 63))
        for fi11 := 0; fi11 < 5; fi11++ {
            v10 += p2
            v12 := GenS2{F0: 178, F1: "\n\t", F2: 854}
            v10 += (p2 % (((fi11 ^ 636)) | 1))
            v13 := GenS3{F0: false}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi4
        _ = v10
    }
    v14 := []int{0, p2, p2, p2}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    p1 += -32768
    switch 2147483647 {
    case 376:
        fmt.Println(p0, p2, p1)
    case 633:
        switch -318 {
        case 9223372036854775807:
            p1 -= 1
        case 128:
            fmt.Println(p1)
        }
    case 42:
        switch ("a" + "hello") {
        case "hello":
            p0 -= p0
        case "foo bar":
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        case "a":
            fmt.Println(p0, p2)
        default:
            p0 += p0
        }
    default:
        switch (p2 >> ((p0) & 63)) {
        case 4294967295:
            v15 := GenS2{F0: 4294967295, F1: "foo bar", F2: 0}
            _ = v15
        case 128:
            if len(v14) > 0 {
                v14[0] = p2
            }
        case 32768:
            switch "foo bar" {
            case "":
                var v16 uint16 = 65535
                _ = v16
            }
        default:
            if len(v14) > 0 {
                v14[0] = (^(p2))
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p1
}
func genFn1(p0 int64) bool {
    v17 := []int{}
    var v18 string = "a"
    _ = p0
    _ = v17
    _ = v18
    return (0 > 283)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "a"
    v3 := make(map[uint32]string)
    var v19 int = 128
    v20 := GenS2{F0: 42, F1: v2, F2: -2147483648}
    v3[2] = v2
    for fi21 := 0; fi21 < 1; fi21++ {
        switch (v2 + v2) {
        case "foo bar":
            if false {
                fmt.Println(v19, v2)
                v19 = fi21
            } else {
                if true {
                    fmt.Println(fi21, fi21)
                    fmt.Println(fi21)
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if true {
                            v23 := []int{fi22, (-(v19))}
                            v19 += v19
                            if ((fi22 >> ((2) & 63)) > 42) {
                                v24 := GenS2{F0: 4294967295, F1: ((((v2 + v2) + "a") + "foo bar") + v2), F2: -2147483648}
                                _ = v24
                            }
                            _ = v23
                        } else {
                            var v25 string = v2
                            v2 = "foo bar"
                            _ = v25
                        }
                        v26 := GenS3{F0: (1 >= 0)}
                        v27 := []int{fi22}
                        fmt.Println(fi22, fi22, v2)
                        var v28 int8 = 42
                        if len(v27) > 0 {
                            v27[0] = fi21
                        }
                        var v29 uint = 42
                        var v30 int16 = int16(v29)
                        _ = fi22
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                v31 := []int{(v19 + v19), (v19 >> ((4294967295) & 63)), v19}
                v32 := GenS1{F0: 1}
                _ = v31
                _ = v32
            }
        default:
            v19 = 2147483647
        }
        if (65535 < 42) {
            for fi33 := 0; fi33 < 17; fi33++ {
                v34 := []int{}
                delete(v1, 2147483647)
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                _ = fi33
                _ = v34
            }
            var v35 float32 = (((0.1 / 1e10) / 1e10) + (0.1 - (float32(-2.25) + 0.0)))
            _ = v35
        }
        _ = fi21
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v19 += v19
        _ = fi36
    }
    var v37 int32 = 823
    v2 = v2
    for fi38 := 0; fi38 < 7; fi38++ {
        fmt.Println(v37)
        _ = fi38
    }
    fmt.Println(v37, v2, v37)
    var v39 uint64 = 2
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 uint32 = 2147483647
        var v42 int8 = 1
        for fi43 := 0; fi43 < 12; fi43++ {
            v2 = "hello"
            _ = fi43
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    v39 = (^(v39))
    for fi44 := 0; fi44 < 4; fi44++ {
        v45 := GenS1{F0: genFn0(920, 1, -341)}
        _ = fi44
        _ = v45
    }
    v39 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v39
}

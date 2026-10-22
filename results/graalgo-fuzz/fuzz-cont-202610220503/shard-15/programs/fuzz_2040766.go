package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 int16; F2 int64 }
func genFn0(p0 int16, p1 int16) int16 {
    switch 127 {
    case 546:
        var v7 uint = 2
        _ = v7
    default:
        v8 := GenS1{F0: 559, F1: 54, F2: ""}
        _ = v8
    }
    fmt.Println(p0, p1)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int = 42
        _ = fi9
        _ = v10
    }
    p0 = 865
    _ = p0
    _ = p1
    return (p0 + (^(p1)))
}

func main() {
    var v1 bool = (2 >= 18446744073709551615)
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v4 := []int{-2147483648}
    var v5 int64 = -1
    v6 := []uint32{1, 0}
    switch ("\n\t" + ("a" + "\n\t")) {
    case "":
        v5 = 370
    case "a":
        v5 = (v5 | 404)
    default:
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint = 65535
            v13 := GenS2{F0: genFn0(-32768, 42)}
            var v14 bool = (v1 || true)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    }
    if (("\n\t" + ("foo bar" + "hello")) > ("hello" + ("foo bar" + ("" + ("" + "hello"))))) {
        switch 381 {
        case 32768:
            v15 := GenS3{F0: "hello", F1: 690}
            _ = v15
        }
        var v16 uint16 = 255
        _ = v16
    } else {
        v1 = v1
        v17 := GenS1{F0: 321, F1: 42, F2: "\n\t"}
        _ = v17
    }
    v5 = 1
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint32 = 2147483648
        _ = fi18
        _ = v19
    }
    var v20 int = 128
    for fi21 := 0; fi21 < 8; fi21++ {
        switch fi21 {
        case 128:
            if v1 {
                v22 := GenS4{F0: "hello", F1: genFn0(1000, 1000), F2: (-(v5))}
                var v23 string = "a"
                for fi24 := 0; fi24 < 12; fi24++ {
                    v5 = -9223372036854775808
                    _ = fi24
                }
                _ = v22
                _ = v23
            } else {
                if (!(1.5 != ((3.14159 - ((float32(1.5) * 1.5) / (0.0 + 1.5))) - 3.14159))) {
                    if v1 {
                        switch 2147483647 {
                        case 4294967295:
                            v1 = (v1 && ((4294967295 != 4294967295) || v1))
                        case -554:
                            for fi25 := 0; fi25 < 14; fi25++ {
                                fmt.Println(fi25, v20, v5)
                                _ = fi25
                            }
                        case 42:
                            var v26 string = ""
                            _ = v26
                        }
                        var v27 int = 298
                        _ = v27
                    } else {
                        var v28 uint64 = 4294967295
                        if true {
                            fmt.Println(v1)
                        }
                        _ = v28
                    }
                    switch "a" {
                    case "a":
                        if (((1.5 * 0.0) + 3.14159) <= 0.0) {
                            fmt.Println(v1)
                            var v29 bool = ("foo bar" > "a")
                            v30 := GenS1{F0: 943, F1: 138, F2: ("a" + ("" + ""))}
                            _ = v29
                            _ = v30
                        }
                    default:
                        v3[""] = "a"
                    }
                    v20 += (-(fi21))
                } else {
                    v20 -= fi21
                    v1 = v1
                }
                fmt.Println(v20)
                if true {
                    if (fi21 == v20) {
                        _ = len(v6)
                    } else {
                        var v31 int8 = 1
                        _ = v31
                    }
                    v32 := GenS4{F0: "a", F1: -262, F2: 561}
                    _ = v32
                }
            }
        case -9223372036854775808:
            fmt.Println(v1, fi21)
        }
        _ = fi21
    }
    fmt.Println(v1, v5)
    var v33 uint64 = 18446744073709551615
    if (-1 != 1) {
        v1 = (1 != 1)
    } else {
        v1 = ((v33 * v33) != 42)
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS4{F0: "foo bar", F1: -32768, F2: v5}
            var v36 bool = (703 != 65535)
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    v5 = int64(v33)
    switch (v20 - v20) {
    case -1:
        if true {
            fmt.Println(v1, v20)
            delete(v2, (v20 ^ -2147483648))
        }
    }
    switch v20 {
    case 42:
        if ((!v1) || ((83 >= 127) && (!(42 > 18446744073709551615)))) {
            v20 = v20
        }
    default:
        v20 = 42
    }
    var v37 uint64 = (v33 + (v33 >> ((1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 string, p1 int64) int8 {
    switch 32767 {
    case 1:
        p1 -= (p1 | 1)
    case 937:
        var v7 int16 = 32767
        _ = v7
    case 32768:
        fmt.Println(p0, p0)
    }
    var v8 uint = 255
    _ = p0
    _ = p1
    _ = v8
    return 15
}
func genFn1(p0 int16, p1 int64, p2 int64) int8 {
    v9 := GenS1{F0: 32768}
    var v10 uint16 = 198
    fmt.Println(v10, p2, p0)
    v10 = (v10 / (((v10 & (^(v10)))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return genFn0("foo bar", (p1 * p1))
}

func main() {
    v1 := []int{0, 172, -1, 4294967295}
    v2 := []string{"foo bar", "", ("hello" + "hello")}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    var v5 uint32 = 4294967295
    v6 := make(map[int]uint32)
    switch -1 {
    case -1:
        var v11 uint = 255
        _ = v11
    case 0:
        if false {
            v12 := GenS1{F0: 968}
            if true {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 uint64 = 44
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v5)
                v15 := GenS1{F0: 128}
                _ = v15
            }
            _ = v12
        } else {
            v5 = (v5 ^ (v5 >> ((52) & 63)))
            v16 := GenS2{F0: 193}
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(fi17, fi17)
                _ = fi17
            }
            _ = v16
        }
    }
    var v18 float32 = 1e10
    if ((0 > 9223372036854775807) && ((0.1 - ((v18 * 1e10) / v18)) != v18)) {
        var v19 int = -344
        _ = v19
    } else {
        v20 := GenS3{F0: ("foo bar" + ("a" + "foo bar"))}
        if (32767 == 32767) {
            var v21 uint = 600
            if (!false) {
                var v22 uint16 = 371
                var v23 uint8 = 2
                _ = v22
                _ = v23
            }
            v5 -= v5
            _ = v21
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            fmt.Println(v18)
            _ = fi24
        }
        _ = v20
    }
    var v25 uint32 = (v5 & (v5 & 42))
    v25 = (v25 - v5)
    var v26 int32 = 42
    v27 := GenS2{F0: 2}
    v1 = append(v1, 520)
    v28 := GenS3{F0: ""}
    var v29 uint32 = v25
    var v30 int8 = 1
    var v31 uint32 = v29
    var v32 int8 = 0
    v26 = (v26 - v26)
    for fi33 := 0; fi33 < 16; fi33++ {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        fmt.Println(v31, v29, v32)
        _ = fi33
    }
    var v34 string = ""
    if (true && true) {
        if (0 > 128) {
            var v35 int32 = (v26 * v26)
            switch v34 {
            case "foo bar":
                if false {
                    v1 = append(v1, 9223372036854775807)
                } else {
                    if false {
                        var v36 int8 = (-(v32))
                        switch -9223372036854775808 {
                        case 128:
                            delete(v4, (("foo bar" + "foo bar") + "hello"))
                        case 1:
                            v32 -= v30
                        }
                        v31 = (v25 ^ (v5 << ((uint(v25)) & 63)))
                        _ = v36
                    } else {
                        v30 = 1
                        var v37 int32 = 1
                        v37 = v35
                        _ = v37
                    }
                }
            case "hello":
                v38 := GenS4{F0: 2, F1: ((v34 + v34) == "foo bar"), F2: 510}
                _ = v38
            case "\n\t":
                if false {
                    v25 = (v31 + (v25 * (v31 * v5)))
                } else {
                    if true {
                        v39 := GenS3{F0: v34}
                        var v40 string = v34
                        v18 -= 3.14159
                        _ = v39
                        _ = v40
                    } else {
                        fmt.Println(v18, v26, v30)
                    }
                    v41 := GenS3{F0: v34}
                    _ = v41
                }
            }
            _ = v35
        }
        v5 -= (v29 / ((42) | 1))
        v42 := GenS5{F0: 242}
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

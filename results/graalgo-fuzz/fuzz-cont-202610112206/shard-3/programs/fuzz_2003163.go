package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 string; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint8, p2 int) int {
    if false {
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 int32 = 0
            p0 = p0
            var v7 int8 = -1
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 9223372036854775807, F1: (p1 >> ((807) & 63)), F2: "foo bar"}
        var v9 uint32 = 4294967295
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 bool = false
            switch "a" {
            case "":
                var v12 uint32 = uint32(p1)
                _ = v12
            case "a":
                var v13 int32 = -26
                _ = v13
            }
            v14 := GenS2{F0: ("" + "hello"), F1: 1, F2: 2}
            _ = fi10
            _ = v11
            _ = v14
        }
        fmt.Println(p1, p2)
    }
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            if (p0 == 255) {
                fmt.Println(p2, p1)
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 float64 = 1e10
                    _ = fi16
                    _ = v17
                }
                if true {
                    v18 := GenS1{F0: -9223372036854775808, F1: p0, F2: (("\n\t" + ("hello" + "foo bar")) + ("a" + "hello"))}
                    _ = v18
                }
            }
            fmt.Println(fi15)
            var v19 int8 = 127
            _ = fi15
            _ = v19
        }
    } else {
        p1 -= 42
    }
    var v20 int64 = 0
    if true {
        var v21 int16 = -1
        _ = v21
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := []int{p2, p2}
            fmt.Println(p1, p1, p0)
            fmt.Println(p1, p0, p1)
            _ = fi22
            _ = v23
        }
        v24 := GenS1{F0: 0, F1: (p1 % ((62) | 1)), F2: ("a" + ((("a" + ("foo bar" + "foo bar")) + "foo bar") + ""))}
        fmt.Println(p0, p0)
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return (p2 * 42)
}

func main() {
    v1 := []string{("a" + "hello")}
    var v2 int32 = -884
    var v3 string = "hello"
    v4 := make(map[int]int)
    var v25 uint16 = 42
    _ = len(v1)
    if ((((1.5 / (1e10 * -2.25)) + float64(float32(0.0))) / -2.25) != float64(((1e10 * 0.0) * -2.25))) {
        v3 = (v3 + "\n\t")
        v3 = ""
    }
    if (258 >= 9223372036854775807) {
        v25 = v25
    }
    delete(v4, genFn0(0, 128, genFn0(255, 128, 600)))
    var v26 int64 = 1
    if (!((float64(float32(1.5)) * float64((-2.25 * float32((float64(1.5) * 0.1))))) >= ((0.1 + 1e10) * -2.25))) {
        var v27 int32 = v2
        _ = v27
    } else {
        fmt.Println(v2, v25)
        v2 = int32(v25)
    }
    switch 1 {
    case -570:
        fmt.Println(v3)
    case 32768:
        if (true && (false || true)) {
            v3 = v3
            switch 1 {
            case 1:
                fmt.Println(v25)
            case 4294967295:
                delete(v4, -288)
            case 128:
                var v28 string = "hello"
                _ = v28
            default:
                v29 := GenS3{F0: (^(v25))}
                _ = v29
            }
            var v30 uint = 4294967295
            _ = v30
        }
    default:
        var v31 uint16 = v25
        _ = v31
    }
    v32 := GenS2{F0: "foo bar", F1: v25, F2: 4294967295}
    if false {
        v2 -= (v2 & -2147483648)
    }
    fmt.Println(v25, v2, v26)
    v3 = v3
    var v33 int32 = (v2 % (((v2 * v2)) | 1))
    _ = len(v1)
    fmt.Println(v33, v33)
    _ = len(v1)
    fmt.Println(v3)
    if ((!(true || false)) && ((v25 * (v25 * 1)) > v25)) {
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 uint64 = 2
            _ = fi34
            _ = v35
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}

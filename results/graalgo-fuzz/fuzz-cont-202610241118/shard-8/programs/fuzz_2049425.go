package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
func genFn0(p0 int8) int {
    for fi5 := 0; fi5 < 2; fi5++ {
        switch "a" {
        case "a":
            fmt.Println(p0, p0)
        }
        _ = fi5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        if (!true) {
            var v7 float64 = ((1e10 * float64((((-2.25 + 1.5) * 1.5) * (1.5 - 0.0)))) - 1.5)
            fmt.Println(p0)
            _ = v7
        } else {
            var v8 int16 = 165
            fmt.Println(v8, p0)
            _ = v8
        }
        switch ("\n\t" + "hello") {
        case "foo bar":
            v9 := []int{(fi6 >> ((255) & 63)), fi6, fi6}
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := []int{}
                v11 = append(v11, (fi6 << ((65535) & 63)))
                _ = fi10
                _ = v11
            }
        case "\n\t":
            if (2147483647 < 42) {
                p0 = p0
                v12 := []int{fi6}
                _ = v12
            }
        default:
            v13 := GenS1{F0: -9223372036854775808, F1: 1}
            _ = v13
        }
        _ = fi6
    }
    v14 := []int{}
    _ = p0
    _ = v14
    return -831
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = (-2.25 * -2.25)
    var v3 uint = 2
    v4 := make(map[string]int)
    v3 -= 436
    if true {
        v3 = v3
        if (-861 == 1) {
            var v15 bool = (51 < 37)
            var v16 uint16 = 0
            _ = v15
            _ = v16
        }
    }
    var v17 uint8 = 255
    switch "hello" {
    case "\n\t":
        switch "hello" {
        case "\n\t":
            var v18 float64 = 0.0
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 18; fi19++ {
                v2 = v2
                v20 := v4["foo bar"]
                v2 = 0.0
                _ = fi19
                _ = v20
            }
        case "a":
            for fi21 := 0; fi21 < 8; fi21++ {
                v1 += v1
                v17 += (v1 << (((v3 / ((v3) | 1))) & 63))
                v1 = (v1 ^ 200)
                _ = fi21
            }
        default:
            if ("a" <= "foo bar") {
                var v22 int = 2147483647
                fmt.Println(v3)
                var v23 uint8 = 0
                _ = v22
                _ = v23
            }
        }
    case "a":
        var v24 int64 = 9223372036854775807
        _ = v24
    }
    v2 -= 0.0
    if false {
        v25 := []int{409, 9223372036854775807, 112}
        v26 := GenS2{F0: 154}
        if false {
            v17 = v1
            v1 += v1
            if (-1 >= -313) {
                fmt.Println(v1, v17)
                var v27 int8 = 100
                _ = v27
            } else {
                var v28 int8 = 1
                _ = v28
            }
        }
        _ = v25
        _ = v26
    } else {
        fmt.Println(v3, v17)
        var v29 bool = true
        _ = v29
    }
    if true {
        if (32767 == genFn0(1)) {
            if ("hello" > (((("\n\t" + ((("foo bar" + (("a" + "foo bar") + "hello")) + "a") + "\n\t")) + "") + "hello") + "\n\t")) {
                v30 := GenS3{F0: (v3 + 4294967295), F1: 0, F2: 568}
                _ = v30
            } else {
                fmt.Println(v17, v2)
                var v31 uint16 = 709
                var v32 uint64 = 18446744073709551615
                _ = v31
                _ = v32
            }
            v33 := GenS2{F0: 4294967295}
            var v34 int = -1
            _ = v33
            _ = v34
        } else {
            var v35 int = 0
            var v36 int = v35
            if true {
                fmt.Println(v36, v2)
            }
            _ = v35
            _ = v36
        }
        v37 := GenS1{F0: -9223372036854775808, F1: 4294967295}
        v38 := GenS2{F0: 212}
        _ = v37
        _ = v38
    }
    v39 := []int{-9223372036854775808, 128, 2147483647, 1}
    fmt.Println(v2)
    v40 := GenS3{F0: (v3 & 177), F1: 735, F2: 1}
    v17 -= 152
    fmt.Println(v3)
    var v41 int8 = 1
    var v42 uint64 = 849
    if true {
        var v43 uint16 = 1
        if ((v2 <= v2) || (1000 != 0)) {
            v41 -= (v41 / (((v41 * -17)) | 1))
            v3 = (v3 * (v3 >> ((v3) & 63)))
            v43 -= (v43 / ((65535) | 1))
        } else {
            fmt.Println(v43)
        }
        _ = v43
    }
    if true {
        v44 := GenS2{F0: v42}
        _ = v44
    } else {
        var v45 int16 = 42
        _ = v45
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}

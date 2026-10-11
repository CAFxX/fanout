package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int, p1 bool) int64 {
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "foo bar":
        var v4 uint32 = 42
        _ = v4
    case "\n\t":
        p1 = ((0.1 / float64(3.14159)) >= (((0.1 * (float64((3.14159 - (0.0 * 1.5))) / 1.0)) * 0.0) + 3.14159))
    }
    v5 := []int{p0}
    switch ("hello" + "foo bar") {
    case "":
        switch (p0 + (p0 / (((p0 + 2147483647)) | 1))) {
        case 9223372036854775807:
            fmt.Println(p1, p1)
        }
    case "\n\t":
        for fi6 := 0; fi6 < 6; fi6++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            p0 = p0
            _ = fi6
        }
    case "foo bar":
        for fi7 := 0; fi7 < 13; fi7++ {
            if ((2 <= 2) && p1) {
                if len(v5) > 0 {
                    v5[0] = 32768
                }
                var v8 int32 = 42
                _ = v8
            } else {
                var v9 bool = true
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v9
            }
            _ = fi7
        }
    }
    if (false && p1) {
        if p1 {
            p0 = p0
        }
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v5
    return 925
}

func main() {
    var v1 int8 = 0
    var v2 int16 = 42
    v3 := make(map[string]uint32)
    var v10 int32 = -191
    switch 4294967295 {
    case 32767:
        switch -329 {
        case -893:
            v2 -= 0
        }
    default:
        v10 -= v10
    }
    v2 -= v2
    for fi11 := 0; fi11 < 11; fi11++ {
        v1 = 42
        _ = fi11
    }
    switch "a" {
    case "\n\t":
        v12 := []int{32768}
        _ = v12
    case "":
        v10 = v10
    }
    v13 := GenS1{F0: 42}
    v1 += v1
    if ((-2.25 * (((1e10 + -2.25) + float32(1.5)) * 0.0)) != 1.5) {
        v1 = v1
        switch "a" {
        case "\n\t":
            if false {
                v2 += v2
                v1 = (v1 + 1)
                var v14 uint64 = 2
                _ = v14
            }
        case "a":
            v15 := []int{int(v1)}
            _ = v15
        case "hello":
            v1 = v1
        }
        var v16 int = -2147483648
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
}

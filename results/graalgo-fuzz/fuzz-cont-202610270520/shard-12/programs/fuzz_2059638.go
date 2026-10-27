package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint16; F1 int8; F2 int }
func genFn0() uint64 {
    if (42 > 42) {
        var v5 uint8 = 16
        switch 128 {
        case -2147483648:
            v6 := GenS1{F0: -1, F1: (v5 % ((v5) | 1))}
            _ = v6
        }
        fmt.Println(v5)
        _ = v5
    }
    if false {
        var v7 bool = ((0 != 775) || true)
        _ = v7
    }
    var v8 float64 = 1.5
    var v9 uint64 = 42
    _ = v8
    _ = v9
    return (v9 * 18446744073709551615)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 491
    var v3 int64 = 0
    var v4 int64 = 1
    v2 -= (v2 * v2)
    if ((!true) || true) {
        var v10 uint16 = 407
        v4 = 0
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                if (-1 != -280) {
                    fmt.Println(fi11, v3, fi11)
                } else {
                    _ = len(v1)
                    var v12 float32 = ((1e10 + (-2.25 * 3.14159)) * 0.1)
                    _ = v12
                }
                _ = fi11
            }
            var v13 uint = 18446744073709551615
            if false {
                var v14 uint64 = 701
                v15 := GenS1{F0: 550, F1: 109}
                _ = v14
                _ = v15
            }
            _ = v13
        }
        _ = v10
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            v17 := GenS2{F0: -1, F1: v4}
            _ = v17
        }
        delete(v1, (("hello" + "hello") + "a"))
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint8 = 1
            _ = fi18
            _ = v19
        }
        _ = fi16
    }
    fmt.Println(v3, v2, v2)
    var v20 uint = 65535
    switch 9223372036854775807 {
    case -630:
        fmt.Println(v3, v4, v4)
    case -2147483648:
        fmt.Println(v4)
    case 4294967295:
        v2 -= -1
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := []int{v2}
        for fi23 := 0; fi23 < 7; fi23++ {
            _ = len(v1)
            switch (("\n\t" + "hello") + "a") {
            case "\n\t":
                var v24 int64 = (v4 * v3)
                v25 := GenS3{F0: 1, F1: 93, F2: v2}
                var v26 bool = true
                _ = v24
                _ = v25
                _ = v26
            case "hello":
                var v27 string = "hello"
                _ = v27
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        if true {
            v2 = (v2 + v2)
            switch -2147483648 {
            case 42:
                for fi29 := 0; fi29 < 18; fi29++ {
                    fmt.Println(fi28, v3)
                    var v30 int8 = -10
                    _ = fi29
                    _ = v30
                }
            case 1:
                var v31 uint16 = 42
                _ = v31
            default:
                fmt.Println(v20, v3, v2)
            }
        } else {
            var v32 uint = v20
            var v33 int = 32768
            for fi34 := 0; fi34 < 10; fi34++ {
                v3 -= v3
                v35 := []int{(^(v2)), 2147483647, (v33 * 0), v33}
                v1["hello"] = 850
                v35 = append(v35, (fi28 * (fi28 ^ (-(fi28)))))
                _ = fi34
                _ = v35
            }
            _ = v32
            _ = v33
        }
        _ = fi28
    }
    fmt.Println(v4)
    switch "a" {
    case "hello":
        var v36 int32 = 563
        _ = v36
    case "\n\t":
        v20 = (v20 % ((v20) | 1))
    case "foo bar":
        v2 = (v2 << ((0) & 63))
    }
    v37 := GenS3{F0: 1, F1: 127, F2: 0}
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 int = (fi38 & fi38)
        _ = fi38
        _ = v39
    }
    v3 = (-(v4))
    var v40 uint64 = 42
    v4 = (v3 ^ (v3 ^ (v4 % ((-9223372036854775808) | 1))))
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 int32 = -1
        _ = fi41
        _ = v42
    }
    var v43 int = (v2 * v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v43
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0() int8 {
    var v7 int64 = 42
    v8 := GenS1{F0: 96, F1: 2, F2: 235}
    var v9 int32 = -632
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "foo bar" {
        case "\n\t":
            v11 := []int{fi10, fi10}
            _ = v11
        case "hello":
            fmt.Println(v9)
        }
        v9 = v9
        var v12 string = (("a" + "\n\t") + ("a" + ("a" + ("a" + (((("foo bar" + ((("foo bar" + "\n\t") + "foo bar") + "a")) + "foo bar") + (("foo bar" + ("foo bar" + ("" + "foo bar"))) + "\n\t")) + "hello")))))
        _ = fi10
        _ = v12
    }
    if false {
        var v13 uint64 = 599
        var v14 float64 = (1.5 * 1e10)
        _ = v13
        _ = v14
    } else {
        v7 = v7
        if true {
            var v15 uint = 676
            v7 = v7
            _ = v15
        }
        var v16 uint64 = 42
        _ = v16
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint8, p1 string, p2 string) int8 {
    if true {
        if (32767 > -1) {
            var v17 uint16 = 2
            _ = v17
        } else {
            fmt.Println(p2)
            fmt.Println(p0, p2)
            p0 -= 2
        }
        if true {
            if false {
                var v18 int64 = 1
                _ = v18
            }
        }
    }
    if (true && false) {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS1{F0: 75, F1: 1, F2: 465}
            var v21 uint = 1
            v21 = 872
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int32 = -2147483648
        p1 = ("\n\t" + p1)
        _ = v22
    } else {
        v23 := GenS1{F0: (p0 & (p0 << ((2) & 63))), F1: 358, F2: 65535}
        p1 = ("foo bar" + (((p1 + "\n\t") + (p2 + (("hello" + (("foo bar" + p1) + p1)) + (p2 + ("hello" + ""))))) + p2))
        v24 := []int{128, 4294967295}
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 string = "hello"
    var v4 string = v3
    v5 := make(map[string]uint32)
    v6 := make(map[int]bool)
    var v25 float32 = float32(-2.25)
    var v26 int32 = 57
    for fi27 := 0; fi27 < 12; fi27++ {
        if (true || (0 <= 440)) {
            v25 = ((v25 * v25) + (float32(0.1) / (v25 + v25)))
            v4 = v3
        }
        var v28 int64 = 9223372036854775807
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 float32 = v25
        v31 := GenS2{F0: 240, F1: 26}
        v32 := []int{(fi29 * 127), (-(fi29)), fi29, fi29}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    if true {
        fmt.Println(v25, v3)
        v26 = 0
        if (149 == 1) {
            v3 = v3
            v26 = (v26 | v26)
            v26 = (v26 / (((v26 - (v26 >> ((18446744073709551615) & 63)))) | 1))
        }
    } else {
        v25 += float32((1e10 * ((float64((v25 / 1.0)) / (((3.14159 * ((((0.1 - 1.5) / (0.0 + 0.1)) + 1.5) - (1e10 * ((0.0 + (0.1 - 1.5)) - 3.14159)))) / 0.1) * 0.1)) + 0.0)))
        v33 := []int{127, -198, 2147483647, 9223372036854775807}
        switch (("\n\t" + v3) + v3) {
        case "hello":
            for fi34 := 0; fi34 < 19; fi34++ {
                fmt.Println(v25)
                v25 += 1.5
                _ = fi34
            }
        case "\n\t":
            var v35 int64 = -9223372036854775808
            _ = v35
        }
        _ = v33
    }
    v36 := v6[128]
    if v36 {
        fmt.Println(v25, v4, v3)
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        v25 = (((1.5 - float32(1e10)) + float32(((1e10 / -2.25) + 0.0))) + (3.14159 * v25))
        delete(v6, fi37)
        _ = fi37
    }
    switch 128 {
    case 32768:
        switch 1 {
        case 1:
            var v38 uint = 0
            _ = v38
        }
    }
    var v39 float32 = -2.25
    var v40 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

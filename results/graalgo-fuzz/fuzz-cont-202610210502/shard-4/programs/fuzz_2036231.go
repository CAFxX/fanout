package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 uint32, p2 uint32) int8 {
    for fi4 := 0; fi4 < 2; fi4++ {
        if (true && (!false)) {
            v5 := []int{}
            _ = v5
        } else {
            v6 := []int{}
            p0 -= 255
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{946, 42, fi7, fi7}
            v9 := GenS1{F0: 2}
            if true {
                p1 = (p1 / (((^(p2))) | 1))
                v10 := GenS2{F0: (fi4 - fi7)}
                _ = v10
            } else {
                switch fi7 {
                case 127:
                    var v11 uint32 = 42
                    _ = v11
                }
                var v12 uint16 = 65535
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        switch (fi4 & (fi4 % ((fi4) | 1))) {
        case -305:
            if (-1 != 42) {
                switch fi4 {
                case 2147483647:
                    var v13 int16 = 693
                    _ = v13
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := []int{392, -2147483648, 2147483647}
                    var v16 float64 = (1e10 / 1.0)
                    switch "\n\t" {
                    case "hello":
                        var v17 int = fi14
                        _ = v17
                    case "foo bar":
                        v15 = append(v15, 0)
                    default:
                        var v18 uint32 = p1
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                p1 = 2
            }
        case -2147483648:
            v19 := GenS2{F0: -680}
            _ = v19
        }
        _ = fi4
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(p1)
        v21 := GenS1{F0: 255}
        p1 = 0
        _ = fi20
        _ = v21
    }
    if true {
        v22 := []int{489}
        fmt.Println(p2, p1)
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return int8(p0)
}
func genFn1() int16 {
    var v23 float64 = 1e10
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 bool = false
        var v26 int64 = 0
        _ = fi24
        _ = v25
        _ = v26
    }
    v23 = v23
    for fi27 := 0; fi27 < 18; fi27++ {
        v28 := []int{0, (fi27 * 32767)}
        var v29 uint16 = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = v23
    return -958
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 1
    var v3 bool = (4294967295 < 42)
    if (!(v3 || (-2.25 >= (((-2.25 * (-2.25 / 0.1)) * (float64((-2.25 / 3.14159)) - ((((0.0 + (0.1 / ((-2.25 * 1e10) - 1e10))) - 1e10) / (1e10 + 0.0)) / 1.0))) + (0.0 * 1e10))))) {
        var v30 int64 = 1
        _ = v30
    } else {
        var v31 float64 = 0.0
        if false {
            var v32 int64 = 9223372036854775807
            _ = v32
        }
        _ = v31
    }
    v33 := GenS1{F0: 540}
    var v34 string = "hello"
    fmt.Println(v1, v1, v1)
    fmt.Println(v34, v1)
    var v35 int = 1
    if (-129 < genFn1()) {
        v36 := []int{}
        _ = v36
    }
    var v37 int8 = 1
    for fi38 := 0; fi38 < 11; fi38++ {
        for fi39 := 0; fi39 < 14; fi39++ {
            fmt.Println(v3, v35)
            v34 = v34
            _ = fi39
        }
        switch (v35 & fi38) {
        case 72:
            v1 += v1
        }
        _ = fi38
    }
    v35 = -2147483648
    var v40 string = v34
    fmt.Println(v1)
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := []int{(fi41 >> ((255) & 63)), (^(fi41)), fi41, fi41}
        v2 += (v2 ^ v2)
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v40
}

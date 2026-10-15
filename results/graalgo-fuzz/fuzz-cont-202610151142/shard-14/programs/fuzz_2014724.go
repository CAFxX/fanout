package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int }
func genFn0() int {
    switch 4294967295 {
    case 2147483647:
        v7 := []int{4294967295}
        _ = v7
    case 32767:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int64 = 9223372036854775807
            switch (-(fi8)) {
            case 127:
                switch fi8 {
                case 2147483647:
                    var v10 int8 = 0
                    _ = v10
                case 128:
                    v9 = v9
                default:
                    v9 = int64(fi8)
                }
            case 746:
                v9 -= int64(fi8)
            default:
                v9 = v9
            }
            _ = fi8
            _ = v9
        }
    }
    if (2 > 1) {
        v11 := GenS1{F0: 127}
        _ = v11
    }
    v12 := GenS2{F0: true, F1: "\n\t"}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = 127
        if true {
            fmt.Println(v14)
            v15 := []int{fi13, (fi13 + fi13)}
            if len(v15) > 0 {
                v15[0] = (fi13 - (^(fi13)))
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    _ = v12
    return 0
}
func genFn1(p0 int64, p1 int8, p2 int32) uint32 {
    switch (("\n\t" + (("\n\t" + "") + "a")) + (("a" + "") + "\n\t")) {
    case "foo bar":
        v16 := GenS3{F0: -922}
        _ = v16
    case "a":
        switch "foo bar" {
        case "":
            if false {
                var v17 string = "a"
                v18 := []int{}
                _ = v17
                _ = v18
            }
        }
    case "hello":
        v19 := []int{346, 32767}
        _ = v19
    }
    p1 = 100
    switch -285 {
    case 2147483647:
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(p1, p1, p0)
            p0 += -419
            _ = fi20
        }
    }
    p0 += 9223372036854775807
    if ((p1 & -128) <= int8(p2)) {
        p1 += p1
        for fi21 := 0; fi21 < 20; fi21++ {
            p0 = (p0 - 1)
            var v22 int = int(p1)
            _ = fi21
            _ = v22
        }
    } else {
        var v23 int64 = (p0 & int64(p1))
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 564
    v3 := make(map[string]bool)
    var v4 int32 = 640
    var v5 int32 = -2147483648
    var v6 uint16 = 761
    switch "foo bar" {
    case "":
        if false {
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS3{F0: fi24}
                _ = fi24
                _ = v25
            }
            v26 := []int{32768, 673}
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 int = (-(fi27))
                if true {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        v30 := v1[(v28 >> ((2) & 63))]
                        if ((!false) || (2147483647 < 0)) {
                            fmt.Println(v5, v30)
                            var v31 uint = 223
                            v32 := v3[("\n\t" + "hello")]
                            _ = v31
                            _ = v32
                        } else {
                            _ = len(v26)
                            var v33 string = "hello"
                            _ = v33
                        }
                        _ = fi29
                        _ = v30
                    }
                }
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    }
    switch 878 {
    case 32767:
        var v34 string = "a"
        _ = v34
    case 2147483647:
        fmt.Println(v4)
    }
    var v35 float64 = ((0.1 - 0.1) - ((float64((1.5 / 1.5)) - float64(float32(float64((1e10 / (1.5 * -2.25)))))) / 1.0))
    v35 += v35
    var v36 int16 = -32768
    v37 := GenS2{F0: false, F1: "hello"}
    v36 += v36
    v5 -= v4
    var v38 bool = (v36 <= (^(v36)))
    var v39 float64 = v35
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}

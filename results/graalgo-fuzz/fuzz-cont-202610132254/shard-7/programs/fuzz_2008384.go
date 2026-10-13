package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0() uint16 {
    var v6 int = 127
    fmt.Println(v6)
    var v7 uint = 599
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 string) int8 {
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 bool = true
        v10 := GenS1{F0: ((p1 + "foo bar") + p1), F1: p1}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            if true {
                switch (p1 + p1) {
                case "a":
                    if true {
                        v13 := []int{}
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = v13
                    } else {
                        if false {
                            v14 := GenS1{F0: p1, F1: p1}
                            _ = v14
                        }
                    }
                case "":
                    var v15 uint16 = 0
                    _ = v15
                default:
                    var v16 int16 = 76
                    _ = v16
                }
            }
            v17 := []int{int(p0), 9223372036854775807, -9223372036854775808}
            p1 = (("\n\t" + ("foo bar" + p1)) + "\n\t")
            _ = fi12
            _ = v17
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 bool = false
            v20 := []int{fi18}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    return -19
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 int8 = -27
    v4 := make(map[string]uint32)
    var v5 float64 = float64(1.5)
    for fi21 := 0; fi21 < 8; fi21++ {
        if (false || false) {
            var v22 uint16 = 32767
            _ = v22
        } else {
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(fi23, fi21, fi21)
                v24 := []int{fi21}
                var v25 float32 = 3.14159
                var v26 float64 = v5
                switch ("foo bar" + ("hello" + ("hello" + "a"))) {
                case "hello":
                    var v27 int8 = (v3 << ((18446744073709551615) & 63))
                    _ = v27
                case "a":
                    v25 = float32((((v5 / v5) * v5) * -2.25))
                }
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        _ = len(v4)
        var v28 int64 = 0
        _ = fi21
        _ = v28
    }
    fmt.Println(v5)
    var v29 uint16 = genFn0()
    if false {
        v29 = genFn0()
    }
    v5 = (0.1 * (float64(1e10) * 0.0))
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 float32 = (3.14159 + (0.0 * 1e10))
        for fi32 := 0; fi32 < 8; fi32++ {
            fmt.Println(fi32, v5)
            var v33 bool = true
            v34 := GenS1{F0: "a", F1: "foo bar"}
            v35 := GenS1{F0: "foo bar", F1: "\n\t"}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi30
        _ = v31
    }
    var v36 int16 = 1
    var v37 string = "hello"
    for fi38 := 0; fi38 < 10; fi38++ {
        v29 += (^(v29))
        _ = fi38
    }
    switch 4294967295 {
    case 1:
        var v39 string = ("a" + (v37 + ((v37 + ("\n\t" + v37)) + "foo bar")))
        _ = v39
    case -984:
        var v40 int16 = (v36 << ((2) & 63))
        _ = v40
    case 2147483647:
        v36 += 1
    }
    v3 -= v3
    var v41 int = -9
    var v42 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}

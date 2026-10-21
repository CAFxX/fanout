package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int = 9223372036854775807
    var v2 float32 = 3.14159
    var v3 int32 = 42
    v4 := make(map[string]int)
    var v5 int64 = -9223372036854775808
    v4[("foo bar" + "a")] = v1
    for fi6 := 0; fi6 < 20; fi6++ {
        v3 += (v3 - (v3 << ((18446744073709551615) & 63)))
        _ = fi6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v1, v2)
    for fi9 := 0; fi9 < 15; fi9++ {
        v3 = (-(v3))
        var v10 uint32 = 985
        v11 := GenS1{F0: 1, F1: 2147483647}
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 = (^(v1))
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint64 = 2
        var v14 uint64 = (v13 << ((255) & 63))
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                v1 = fi15
                var v16 uint = 2
                var v17 float32 = v2
                v18 := []int{}
                var v19 int8 = 100
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 bool = true
                v2 = (float32(3.14159) / (-2.25 * float32((3.14159 / (float64(1.5) + 1.5)))))
                v21 := GenS2{F0: 0}
                _ = v20
                _ = v21
            }
            _ = fi15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v22 uint = 4294967295
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint8 = 42
        var v25 int64 = -565
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        if true {
            v3 = 770
            for fi27 := 0; fi27 < 17; fi27++ {
                if (255 != 65535) {
                    v28 := v4["foo bar"]
                    v29 := []int{0, v1, 27}
                    _ = v28
                    _ = v29
                } else {
                    switch fi26 {
                    case 42:
                        v2 += ((v2 - v2) - (v2 / v2))
                    }
                }
                switch ((((("hello" + "a") + ("foo bar" + ("foo bar" + "a"))) + "a") + (("\n\t" + "\n\t") + "hello")) + ("a" + "a")) {
                case "\n\t":
                    if (!false) {
                        if (true && (true || (false && (319 > (v22 & 2))))) {
                            var v30 uint = (v22 - v22)
                            v31 := GenS1{F0: 255, F1: 0}
                            _ = v30
                            _ = v31
                        } else {
                            if (((!(!(true || (889 > v22)))) && false) && false) {
                                v32 := GenS3{F0: 42, F1: 65535, F2: ((0.1 / v2) == (v2 + float32(float64(-2.25))))}
                                _ = v32
                            } else {
                                v2 = float32(1.5)
                            }
                            var v33 int8 = -100
                            v34 := v4["foo bar"]
                            _ = v33
                            _ = v34
                        }
                        var v35 float32 = 0.1
                        var v36 uint64 = 613
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi27
            }
            var v37 int64 = 42
            _ = v37
        }
        fmt.Println(fi26, v22)
        _ = fi26
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        v39 := GenS2{F0: 42}
        if true {
            v5 = v5
            v40 := GenS4{F0: 42}
            _ = v40
        }
        fmt.Println(v22, v22)
        _ = fi38
        _ = v39
    }
    v2 = v2
    v41 := []int{v1, (v1 >> ((421) & 63)), 4294967295}
    var v42 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}

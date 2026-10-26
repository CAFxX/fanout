package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 uint32, p2 bool) bool {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint16 = 1
        p1 += 0
        p2 = (fi6 <= (fi6 * int(p0)))
        _ = fi6
        _ = v7
    }
    var v8 uint = p0
    switch "\n\t" {
    case "":
        if true {
            v9 := []int{32767, 32767, 0}
            _ = v9
        }
    }
    if (0 >= 1) {
        var v10 uint8 = 53
        _ = v10
    }
    var v11 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return (2147483647 <= 0)
}

func main() {
    var v1 bool = (42 < 255)
    v2 := make(map[string]string)
    var v3 float64 = float64(0.0)
    v4 := []int{127, -1}
    var v5 int32 = -543
    v12 := GenS1{F0: 1}
    if (1 <= 914) {
        if (1 == 749) {
            var v13 float32 = 0.0
            fmt.Println(v5, v5, v1)
            _ = len(v4)
            _ = v13
        }
        switch "\n\t" {
        case "a":
            var v14 float32 = ((0.1 + -2.25) + ((float32((float64(float32((v3 - (v3 / v3)))) * (0.1 + v3))) * (float32((v3 / 1.0)) + 1.5)) - 1.5))
            _ = v14
        case "hello":
            fmt.Println(v3, v1, v1)
        default:
            v5 -= (v5 & v5)
        }
    }
    v15 := GenS1{F0: -1}
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 float32 = ((0.1 + 3.14159) - 3.14159)
        _ = fi16
        _ = v17
    }
    v5 = (v5 * v5)
    fmt.Println(v1)
    var v18 uint16 = 1
    var v19 uint32 = 0
    v1 = (!v1)
    var v20 string = "foo bar"
    fmt.Println(v18)
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            switch v20 {
            case "foo bar":
                if (uint8(v19) == 255) {
                    v20 = ("hello" + "")
                    if ((v18 >> ((467) & 63)) <= v18) {
                        if (2 < 502) {
                            if v1 {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        v23 := GenS2{F0: 2}
                        fmt.Println(fi21, v18)
                        _ = v23
                    }
                    var v24 uint = 255
                    _ = v24
                } else {
                    switch v20 {
                    case "hello":
                        if (fi21 == -2147483648) {
                            var v25 uint16 = (^(v18))
                            var v26 bool = (true && v1)
                            _ = v25
                            _ = v26
                        }
                        var v27 int64 = 0
                        if v1 {
                            v28 := GenS1{F0: v27}
                            _ = v28
                        }
                        _ = v27
                    case "":
                        var v29 uint = 0
                        _ = v29
                    case "\n\t":
                        var v30 uint16 = v18
                        _ = v30
                    }
                }
            }
            v18 -= v18
            if true {
                if (42 <= v5) {
                    _ = len(v2)
                    v3 = 1e10
                }
            } else {
                var v31 uint16 = v18
                _ = v31
            }
            _ = fi22
        }
        switch (fi21 - (fi21 * (fi21 / (((fi21 - fi21)) | 1)))) {
        case 4294967295:
            var v32 bool = v1
            _ = v32
        case 32767:
            var v33 uint64 = 18446744073709551615
            _ = v33
        default:
            switch ((v20 + "\n\t") + "a") {
            case "foo bar":
                v34 := GenS2{F0: 1}
                _ = v34
            default:
                v3 = ((v3 / (v3 / v3)) + ((v3 * (v3 * v3)) + v3))
            }
        }
        fmt.Println(v5, v19)
        _ = fi21
    }
    delete(v2, ((v20 + v20) + (v20 + ("\n\t" + v20))))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

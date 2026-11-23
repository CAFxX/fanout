package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint64, p1 int32) int {
    fmt.Println(p1)
    var v6 int64 = -430
    var v7 uint16 = 32767
    p1 -= (p1 | p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 973
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = 100
    var v3 uint16 = 825
    v4 := make(map[int]string)
    var v5 int32 = -891
    var v8 uint16 = v3
    if (-1 != 42) {
        delete(v4, 127)
    }
    var v9 int16 = 32767
    var v10 uint = 4294967295
    if true {
        fmt.Println(v9, v1, v3)
        if (9223372036854775807 <= -21) {
            v1 = ((v1 + (v1 - v1)) * v1)
            if false {
                if (("" != "hello") || false) {
                    if (0 == 18446744073709551615) {
                        if ((true && ("hello" > "foo bar")) && ("foo bar" <= (((("\n\t" + ("\n\t" + "hello")) + ("hello" + "\n\t")) + ((((("hello" + "foo bar") + (("hello" + "\n\t") + "foo bar")) + "foo bar") + "") + "\n\t")) + "a"))) {
                            if ((v3 ^ v3) != v3) {
                                v5 = v5
                            }
                            for fi11 := 0; fi11 < 7; fi11++ {
                                for fi12 := 0; fi12 < 1; fi12++ {
                                    var v13 uint = 4294967295
                                    var v14 float32 = (v1 / float32(0.1))
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                                v15 := GenS1{F0: 4294967295, F1: false}
                                v1 = v1
                                _ = fi11
                                _ = v15
                            }
                        } else {
                            var v16 int64 = -1
                            _ = v16
                        }
                    }
                    var v17 int32 = (v5 * v5)
                    v18 := GenS1{F0: -1, F1: ("foo bar" > "a")}
                    _ = v17
                    _ = v18
                } else {
                    if true {
                        v19 := GenS1{F0: 515, F1: (false && false)}
                        _ = v19
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    switch -314 {
                    case 128:
                        var v20 uint64 = 674
                        _ = v20
                    case 1:
                        var v21 uint16 = 1
                        _ = v21
                    case 42:
                        var v22 int16 = v9
                        _ = v22
                    }
                case "\n\t":
                    var v23 int = 626
                    _ = v23
                case "a":
                    for fi24 := 0; fi24 < 20; fi24++ {
                        v10 -= 63
                        var v25 int32 = 822
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 uint64 = 42
                _ = v26
            }
        } else {
            var v27 uint64 = 18446744073709551615
            fmt.Println(v5, v1)
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 float64 = 0.1
                v30 := GenS1{F0: fi28, F1: true}
                var v31 int16 = v9
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = v27
        }
        if ((32767 != -348) && (((-2.25 + v1) - (v1 + 1e10)) < v1)) {
            v32 := GenS1{F0: 0, F1: true}
            var v33 string = ((("hello" + (((("a" + "a") + "a") + ((("foo bar" + ("foo bar" + "a")) + ("hello" + "a")) + ((("hello" + "hello") + "\n\t") + "hello"))) + "hello")) + "") + "hello")
            fmt.Println(v8)
            _ = v32
            _ = v33
        } else {
            v34 := []int{-875, 4294967295, genFn0(4294967295, 940), 4294967295}
            var v35 int64 = 9223372036854775807
            v36 := GenS2{F0: 32767}
            _ = v34
            _ = v35
            _ = v36
        }
    } else {
        if true {
            v37 := GenS1{F0: -468, F1: true}
            v38 := GenS1{F0: 128, F1: ("hello" <= "a")}
            _ = v37
            _ = v38
        }
    }
    v39 := GenS2{F0: 4294967295}
    var v40 uint8 = 128
    var v41 uint16 = v8
    var v42 float32 = float32(((1.5 + float64(float32(0.0))) / (-2.25 * ((1e10 - 1.5) * 1e10))))
    var v43 uint32 = 4294967295
    for fi44 := 0; fi44 < 20; fi44++ {
        v40 = uint8(v43)
        v45 := GenS3{F0: uint(v3), F1: v10}
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}

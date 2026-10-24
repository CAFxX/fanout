package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
func genFn0() int {
    var v6 uint64 = 42
    v6 += 42
    switch "a" {
    case "foo bar":
        v7 := []int{}
        _ = v7
    }
    v8 := GenS1{F0: 60, F1: (-1 <= -2147483648), F2: 42}
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := []string{("" + "a"), "hello", (("foo bar" + "a") + "a"), "a"}
    var v2 uint64 = 18446744073709551615
    v3 := []int{-2147483648, -1, 2147483647}
    var v4 uint = 65535
    var v5 uint = (v4 - (v4 * v4))
    switch 0 {
    case -689:
        v5 += v4
    case 41:
        var v9 uint = v5
        _ = v9
    case 32768:
        if false {
            var v10 bool = false
            _ = v10
        }
    }
    fmt.Println(v4, v4, v5)
    if (1 >= 444) {
        v3 = append(v3, -742)
        if false {
            if (float64(1.5) < 3.14159) {
                fmt.Println(v4)
                fmt.Println(v5, v5)
                for fi11 := 0; fi11 < 19; fi11++ {
                    v5 += (v5 / (((v5 / ((0) | 1))) | 1))
                    for fi12 := 0; fi12 < 10; fi12++ {
                        if true {
                            v1 = append(v1, ("\n\t" + ("hello" + "hello")))
                            v2 += (^(v2))
                            var v13 int64 = 724
                            var v14 uint64 = v2
                            var v15 uint = (^(v5))
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        _ = len(v3)
                        var v16 uint64 = 4294967295
                        _ = fi12
                        _ = v16
                    }
                    var v17 uint = (v5 << (((^(v4))) & 63))
                    _ = fi11
                    _ = v17
                }
            }
            var v18 int32 = 0
            _ = v18
        } else {
            v3 = append(v3, 835)
        }
    }
    _ = len(v3)
    v5 -= v4
    for fi19 := 0; fi19 < 3; fi19++ {
        v1 = append(v1, ("foo bar" + ""))
        var v20 bool = true
        if false {
            var v21 uint64 = v2
            for fi22 := 0; fi22 < 7; fi22++ {
                v4 = 65535
                if len(v1) > 0 {
                    v1[0] = ((("foo bar" + "foo bar") + ("a" + ("\n\t" + ("a" + ("a" + ("\n\t" + "a")))))) + "hello")
                }
                switch ("a" + "a") {
                case "":
                    v2 = v21
                default:
                    if v20 {
                        var v23 string = "hello"
                        _ = v23
                    }
                }
                if v20 {
                    var v24 int = fi19
                    fmt.Println(v4)
                    _ = v24
                } else {
                    var v25 uint = (v5 - (v4 ^ (v4 / (((v4 ^ 4294967295)) | 1))))
                    _ = v25
                }
                if v20 {
                    v26 := GenS1{F0: 127, F1: ((fi22 ^ fi22) == (fi19 * 103)), F2: 697}
                    _ = v26
                } else {
                    v2 += 42
                    fmt.Println(v4, fi19)
                }
                _ = fi22
            }
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    var v27 int16 = 1000
    if (!(float64(0.1) < 0.1)) {
        _ = len(v3)
        v2 += v2
    } else {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v28 bool = (0 == 9223372036854775807)
            v1 = append(v1, "a")
            _ = v28
        } else {
            _ = len(v1)
        }
        var v29 int8 = 0
        fmt.Println(v5, v4)
        _ = v29
    }
    if (false && false) {
        v4 -= v5
        var v30 int = 0
        var v31 string = ("\n\t" + "")
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 bool = true
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 uint64 = (v2 | v2)
        _ = len(v1)
        var v36 uint16 = 810
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

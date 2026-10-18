package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 bool, p2 uint) int32 {
    var v4 uint = 255
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{(p0 * fi5), fi5, (p0 ^ (p0 << ((v4) & 63)))}
        var v7 uint = 563
        if false {
            if false {
                var v8 int16 = -32768
                _ = v8
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{(fi9 & -564)}
        _ = len(v10)
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 0}
            switch (("foo bar" + "foo bar") + "a") {
            case "foo bar":
                switch "a" {
                case "":
                    v13 := GenS1{F0: 69}
                    _ = v13
                default:
                    var v14 uint16 = 65535
                    _ = v14
                }
            case "\n\t":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case "a":
                if p1 {
                    if true {
                        var v15 uint16 = 32767
                        v16 := GenS1{F0: 0}
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(p2)
                    if ("\n\t" != "") {
                        _ = len(v10)
                    } else {
                        var v17 uint64 = 42
                        _ = v17
                    }
                    if (1 == -960) {
                        if p1 {
                            v18 := GenS1{F0: -128}
                            p1 = false
                            switch "\n\t" {
                            case "":
                                p1 = p1
                            default:
                                fmt.Println(fi9, p1, p2)
                            }
                            _ = v18
                        }
                    } else {
                        if false {
                            switch (fi11 + fi11) {
                            case 4294967295:
                                _ = len(v10)
                            }
                            v4 = (v4 | (v4 + v4))
                            fmt.Println(p1, fi9, p1)
                            v19 := GenS2{F0: p1, F1: 0}
                            var v20 uint32 = 65535
                            _ = v19
                            _ = v20
                        }
                        if p1 {
                            if len(v10) > 0 {
                                v10[0] = (p0 >> ((p2) & 63))
                            }
                            p0 = (fi9 ^ fi9)
                            p2 -= p2
                        } else {
                            v4 -= 18446744073709551615
                            var v21 int64 = 1
                            p1 = ("hello" <= "foo bar")
                            _ = v21
                        }
                        v4 = v4
                    }
                    v4 += 327
                }
            default:
                v10 = append(v10, (fi9 / ((fi9) | 1)))
            }
            v4 += 42
            fmt.Println(v4, p0, p2)
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v22 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v4
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 bool = true
    fmt.Println(v2)
    v23 := []int{42}
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint32 = 4294967295
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int8 = 100
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    delete(v1, "\n\t")
    fmt.Println(v3)
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    var v28 bool = true
    var v29 bool = false
    for fi30 := 0; fi30 < 9; fi30++ {
        v31 := GenS2{F0: true, F1: 175}
        v29 = false
        if v29 {
            v28 = true
            v32 := GenS1{F0: -126}
            for fi33 := 0; fi33 < 16; fi33++ {
                switch ("\n\t" + ("" + "a")) {
                case "\n\t":
                    var v34 int16 = 1
                    _ = v34
                case "foo bar":
                    v35 := GenS1{F0: -1}
                    var v36 uint8 = 2
                    fmt.Println(v28)
                    _ = v35
                    _ = v36
                }
                v37 := GenS3{F0: -2147483648}
                v23 = append(v23, -1)
                _ = fi33
                _ = v37
            }
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 float64 = (3.14159 - float64(1.5))
        var v40 int32 = genFn0(fi38, v28, 2)
        _ = fi38
        _ = v39
        _ = v40
    }
    v28 = true
    switch "" {
    case "foo bar":
        var v41 int = -354
        _ = v41
    case "":
        v3 = v28
    case "hello":
        for fi42 := 0; fi42 < 5; fi42++ {
            v2 += (v2 << ((2) & 63))
            _ = fi42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
}

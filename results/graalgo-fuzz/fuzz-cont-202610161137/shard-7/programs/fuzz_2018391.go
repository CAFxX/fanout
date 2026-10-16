package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 int; F2 string }
func genFn0(p0 int64, p1 uint64, p2 string) bool {
    var v7 int = 1
    if true {
        var v8 uint64 = (p1 | p1)
        if (true || (((true && true) || (-1 <= 1000)) && false)) {
            p0 = (-(p0))
            fmt.Println(v8)
            if true {
                v9 := GenS1{F0: -95, F1: 65535}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 13; fi10++ {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v12 := GenS1{F0: 18, F1: 42}
                        v13 := []int{fi11}
                        switch 127 {
                        case 4294967295:
                            switch (^(fi11)) {
                            case 128:
                                var v14 int = 2147483647
                                _ = v14
                            default:
                                v7 += -197
                                v13 = append(v13, 60)
                            }
                        }
                        v13 = append(v13, (fi10 * (fi10 >> ((908) & 63))))
                        if false {
                            fmt.Println(v8, fi11)
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    for fi15 := 0; fi15 < 5; fi15++ {
                        p1 -= p1
                        _ = fi15
                    }
                    _ = fi10
                }
                var v16 int16 = int16(p1)
                _ = v16
            }
        } else {
            v17 := GenS2{F0: v7, F1: v7, F2: (p2 + p2)}
            _ = v17
        }
        _ = v8
    } else {
        var v18 int16 = 247
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 int16 = 42
                v22 := []int{}
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 int64 = p0
            switch p2 {
            case "\n\t":
                if ((((("a" + p2) + ("hello" + p2)) + p2) + "") != (p2 + p2)) {
                    fmt.Println(fi19, v18)
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 int8 = int8(v7)
                        _ = fi24
                        _ = v25
                    }
                    switch (v7 & (-(fi19))) {
                    case 0:
                        p0 -= (p0 + 42)
                    case 32768:
                        p1 = (p1 % ((42) | 1))
                    case -854:
                        var v26 int = (fi19 ^ v7)
                        _ = v26
                    }
                    switch fi19 {
                    case -9223372036854775808:
                        var v27 int64 = 1
                        _ = v27
                    default:
                        v23 += v23
                    }
                }
            case "foo bar":
                fmt.Println(v7, v7, v7)
            }
            _ = fi19
            _ = v23
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 string = p2
            var v30 uint16 = 753
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v18
    }
    var v31 uint = 4294967295
    var v32 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    _ = v7
    return (v7 < 0)
}

func main() {
    var v1 uint16 = 409
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 uint8 = 127
    v5 := make(map[string]string)
    var v6 string = "hello"
    v33 := GenS2{F0: -2147483648, F1: -1, F2: "hello"}
    v4 -= v4
    v1 = v1
    var v34 bool = ((false && true) && true)
    v1 -= (v1 >> ((645) & 63))
    delete(v5, "foo bar")
    v35 := GenS2{F0: 186, F1: 829, F2: ((v6 + (v6 + "")) + "foo bar")}
    var v36 int = 128
    for fi37 := 0; fi37 < 1; fi37++ {
        for fi38 := 0; fi38 < 13; fi38++ {
            v34 = (!v34)
            switch v6 {
            case "\n\t":
                fmt.Println(fi37, v34)
            case "hello":
                var v39 uint64 = 715
                _ = v39
            }
            _ = fi38
        }
        _ = fi37
    }
    var v40 int16 = -1
    _ = len(v3)
    var v41 bool = v34
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}

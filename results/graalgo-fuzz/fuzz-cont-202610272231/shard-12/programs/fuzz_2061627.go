package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint; F1 string; F2 bool }

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 60
    var v3 int64 = 9223372036854775807
    var v4 bool = (!(3.14159 >= float32(float64((3.14159 * (0.1 / 1.5))))))
    v5 := make(map[string]int)
    switch 42 {
    case -9223372036854775808:
        if ("foo bar" != "foo bar") {
            if v4 {
                var v6 uint32 = (v1 << ((2) & 63))
                var v7 uint = 554
                _ = v6
                _ = v7
            } else {
                var v8 string = ""
                fmt.Println(v2, v2)
                _ = v8
            }
        }
    case 927:
        switch -1 {
        case 42:
            var v9 uint64 = 783
            _ = v9
        case -108:
            fmt.Println(v2, v1)
        case 128:
            if false {
                v10 := []int{42}
                _ = len(v5)
                v4 = (!v4)
                _ = v10
            }
        default:
            v5[("a" + "\n\t")] = -2147483648
        }
    default:
        v1 += v1
    }
    v1 += (v1 / (((v1 & v1)) | 1))
    v1 += (v1 / ((v1) | 1))
    if true {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int32 = 42
            _ = fi11
            _ = v12
        }
    } else {
        fmt.Println(v2, v4)
        v1 += v1
        var v13 int8 = 0
        _ = v13
    }
    v14 := GenS1{F0: 1, F1: 1}
    v15 := v5[("hello" + "foo bar")]
    switch ("foo bar" + "hello") {
    case "a":
        v4 = false
    }
    switch "\n\t" {
    case "foo bar":
        v2 += 1
    default:
        if (v15 > (v15 + v15)) {
            v16 := []int{}
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS2{F0: 0, F1: "", F2: (-1 != 1)}
                var v19 uint = 4294967295
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 uint8 = 42
                    v2 = (v2 / ((v2) | 1))
                    if v4 {
                        fmt.Println(v15, fi17, v1)
                    } else {
                        fmt.Println(v3)
                    }
                    _ = fi20
                    _ = v21
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v16
        } else {
            v3 += v3
            v1 += v1
            var v22 uint64 = 42
            _ = v22
        }
    }
    v15 += (v15 % ((2147483647) | 1))
    var v23 uint = 255
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := v5["a"]
        switch "foo bar" {
        case "hello":
            var v26 int = v15
            _ = v26
        default:
            if v4 {
                var v27 float64 = 3.14159
                _ = v27
            } else {
                delete(v5, "a")
            }
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            v25 = (fi24 ^ 9223372036854775807)
            v1 -= (v1 ^ 182)
            var v29 int16 = 32767
            _ = fi28
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    switch "foo bar" {
    case "\n\t":
        switch -175 {
        case 42:
            switch "hello" {
            case "":
                var v30 int8 = -128
                _ = v30
            case "a":
                var v31 uint8 = 255
                _ = v31
            case "hello":
                for fi32 := 0; fi32 < 6; fi32++ {
                    _ = len(v5)
                    _ = fi32
                }
            }
        default:
            var v33 bool = ((v23 >= v23) || (9223372036854775807 == v3))
            _ = v33
        }
    default:
        v34 := GenS1{F0: int16(v23), F1: int32(v3)}
        _ = v34
    }
    v35 := []int{v15, v15, (v15 >> ((255) & 63))}
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 int64 = v3
        for fi38 := 0; fi38 < 5; fi38++ {
            v37 -= -725
            v39 := v5["a"]
            var v40 uint16 = v2
            _ = fi38
            _ = v39
            _ = v40
        }
        v15 = int(v3)
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}

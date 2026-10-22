package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0() int16 {
    if ((32767 < -513) || false) {
        var v6 int = 1
        _ = v6
    }
    switch 1 {
    case -2147483648:
        v7 := GenS1{F0: 4294967295}
        _ = v7
    case 1:
        var v8 int16 = 42
        _ = v8
    case 2147483647:
        var v9 int = 4294967295
        _ = v9
    }
    return 813
}

func main() {
    var v1 int = -708
    var v2 float32 = 1e10
    var v3 float64 = -2.25
    v4 := []string{"hello", "hello", "foo bar", "\n\t"}
    v5 := make(map[int]bool)
    if true {
        v10 := GenS1{F0: 18446744073709551615}
        _ = v10
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v5[(v1 | v1)] = (!((!true) && false))
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int64 = 869
        var v13 uint16 = 2
        _ = fi11
        _ = v12
        _ = v13
    }
    switch ("a" + "foo bar") {
    case "a":
        var v14 int64 = -1
        _ = v14
    case "foo bar":
        var v15 uint8 = 75
        _ = v15
    case "\n\t":
        var v16 uint8 = 2
        _ = v16
    }
    v2 -= v2
    var v17 int8 = 100
    for fi18 := 0; fi18 < 17; fi18++ {
        switch v1 {
        case -658:
            fmt.Println(v3, v1, v1)
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            _ = len(v4)
            var v20 string = ("foo bar" + ("foo bar" + ""))
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v21 uint16 = 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 uint8 = 150
    if (((v22 & (^(v22))) <= v22) && false) {
        if false {
            if len(v4) > 0 {
                v4[0] = ((("foo bar" + "hello") + "hello") + ("hello" + "\n\t"))
            }
        } else {
            var v23 int = v1
            _ = len(v4)
            var v24 int16 = 42
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 int64 = 1
            fmt.Println(v2, v2)
            v26 = v26
            _ = fi25
            _ = v26
        }
        v1 += v1
    } else {
        fmt.Println(v2)
        v2 += -2.25
        var v27 string = "hello"
        _ = v27
    }
    switch (((("\n\t" + "") + (("foo bar" + ("a" + "foo bar")) + "a")) + "\n\t") + "") {
    case "foo bar":
        if true {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v28 := GenS2{F0: (v22 & (^(v22))), F1: (v1 / ((int(v21)) | 1))}
                _ = v28
            case "hello":
                fmt.Println(v1, v2, v17)
            case "":
                v29 := GenS3{F0: (^(v22))}
                _ = v29
            }
        } else {
            v30 := GenS1{F0: 1}
            _ = v30
        }
    case "hello":
        if true {
            fmt.Println(v1, v17, v22)
            v4 = append(v4, "\n\t")
            v22 -= (v22 / ((v22) | 1))
        }
    case "a":
        v1 += int(v21)
    default:
        var v31 bool = (1.5 != (v3 + (v3 * 0.0)))
        _ = v31
    }
    var v32 uint64 = 42
    for fi33 := 0; fi33 < 19; fi33++ {
        v32 += (v32 / ((v32) | 1))
        _ = fi33
    }
    if false {
        v22 = 255
        var v34 uint8 = (v22 ^ (v22 % (((^(v22))) | 1)))
        _ = v34
    } else {
        var v35 int16 = 1
        if (42 != (v1 | (v1 ^ (v1 << ((741) & 63))))) {
            var v36 int = (v1 >> ((4294967295) & 63))
            _ = v36
        }
        var v37 uint32 = 65535
        _ = v35
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}

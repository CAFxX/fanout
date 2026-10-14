package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int64, p1 int64, p2 string) string {
    if true {
        v6 := GenS1{F0: (p2 + "foo bar"), F1: -828}
        v7 := GenS1{F0: ((p2 + p2) + ((p2 + "hello") + (("foo bar" + p2) + "a"))), F1: -9223372036854775808}
        var v8 int32 = -2147483648
        _ = v6
        _ = v7
        _ = v8
    }
    switch p2 {
    case "hello":
        v9 := []int{9223372036854775807}
        _ = v9
    case "a":
        v10 := GenS2{F0: 127}
        _ = v10
    case "foo bar":
        switch 4294967295 {
        case 1:
            fmt.Println(p2, p2, p1)
        }
    }
    switch -9223372036854775808 {
    case 127:
        v11 := GenS3{F0: -32768, F1: 32767, F2: false}
        _ = v11
    case 0:
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint32 = 342
            _ = fi12
            _ = v13
        }
    case -9223372036854775808:
        if false {
            p2 = (p2 + "")
        } else {
            if (false && (574 > 747)) {
                p2 = p2
            }
            switch "" {
            case "\n\t":
                var v14 uint32 = 42
                _ = v14
            }
        }
    }
    var v15 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p2
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    v3 := []string{("a" + ("\n\t" + ("" + ("a" + ("foo bar" + "\n\t")))))}
    var v4 int64 = -9223372036854775808
    var v5 int32 = (v2 & (v2 >> ((18446744073709551615) & 63)))
    var v16 uint = 255
    v17 := GenS2{F0: -1}
    v5 += (v5 >> ((4294967295) & 63))
    v18 := GenS1{F0: ("" + "foo bar"), F1: -854}
    fmt.Println(v2, v2, v4)
    if (4294967295 > 4294967295) {
        switch "a" {
        case "":
            v2 = -892
        case "\n\t":
            fmt.Println(v16)
        }
    } else {
        v2 += v2
        fmt.Println(v5, v2)
    }
    v19 := GenS2{F0: 0}
    v3 = append(v3, "foo bar")
    v5 += -316
    v4 = (v4 / ((v4) | 1))
    var v20 int16 = 42
    v1[-1] = -1
    for fi21 := 0; fi21 < 13; fi21++ {
        switch (fi21 & (fi21 - fi21)) {
        case -529:
            if false {
                for fi22 := 0; fi22 < 6; fi22++ {
                    v4 = (v4 % ((-534) | 1))
                    var v23 int = int(v4)
                    fmt.Println(fi21)
                    _ = fi22
                    _ = v23
                }
                var v24 uint16 = 110
                _ = v24
            }
        }
        v2 -= (v5 & v5)
        _ = fi21
    }
    fmt.Println(v2, v20, v2)
    if (9223372036854775807 >= 811) {
        if ((v20 % ((32767) | 1)) != (v20 ^ (v20 >> (((v16 % ((65535) | 1))) & 63)))) {
            v5 = (v5 + (v5 | v5))
            v2 += 2147483647
        } else {
            _ = len(v1)
            v16 += (^(v16))
        }
        v5 = (v5 * 42)
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS1{F0: "", F1: 128}
        _ = fi25
        _ = v26
    }
    if (!(false || true)) {
        fmt.Println(v4)
    } else {
        v2 = (v5 | 935)
        var v27 float32 = -2.25
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}

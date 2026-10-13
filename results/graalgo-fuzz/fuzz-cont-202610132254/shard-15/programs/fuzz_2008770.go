package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint = 2
    v2 := make(map[int]string)
    var v3 uint16 = 2
    var v4 bool = true
    var v5 int = 4294967295
    if v4 {
        v5 = 42
        switch "\n\t" {
        case "\n\t":
            v4 = v4
        case "foo bar":
            switch v5 {
            case 9223372036854775807:
                v5 = v5
            case -530:
                v1 += 18446744073709551615
            }
        }
    }
    switch ("hello" + ("" + "a")) {
    case "foo bar":
        v3 = v3
    case "a":
        fmt.Println(v3, v1)
    default:
        switch "\n\t" {
        case "foo bar":
            if false {
                var v6 uint32 = 2
                v6 += (v6 - 42)
                var v7 int64 = 1
                _ = v6
                _ = v7
            } else {
                if v4 {
                    fmt.Println(v1, v1, v5)
                    switch (-(v5)) {
                    case -2147483648:
                        fmt.Println(v1, v3, v4)
                    case 42:
                        v1 += v1
                    default:
                        v8 := []int{}
                        _ = v8
                    }
                    v3 = (^(v3))
                }
                v9 := GenS1{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: (v3 - (v3 * 2)), F2: "\n\t"}
                _ = v9
            }
        }
    }
    var v10 int16 = 1000
    fmt.Println(v4, v3)
    for fi11 := 0; fi11 < 4; fi11++ {
        v3 = v3
        if true {
            for fi12 := 0; fi12 < 12; fi12++ {
                v4 = v4
                if (0 > 0) {
                    var v13 uint32 = 4294967295
                    switch ((("a" + "foo bar") + (("foo bar" + "foo bar") + "\n\t")) + "foo bar") {
                    case "a":
                        v14 := GenS1{F0: "a", F1: (v3 ^ v3), F2: ("hello" + ("hello" + ("a" + "")))}
                        _ = v14
                    case "":
                        var v15 uint64 = 1
                        _ = v15
                    case "foo bar":
                        var v16 int16 = (v10 / (((^(v10))) | 1))
                        _ = v16
                    }
                    _ = v13
                } else {
                    v17 := []int{}
                    _ = v17
                }
                _ = fi12
            }
        }
        _ = fi11
    }
    v10 += v10
    v18 := GenS2{F0: 521, F1: 0, F2: uint32(v1)}
    v1 += (v1 - v1)
    switch (v5 << ((v1) & 63)) {
    case 0:
        if v4 {
            fmt.Println(v4)
        } else {
            v3 += 1
            v19 := GenS3{F0: 948}
            _ = v19
        }
    case 127:
        if v4 {
            var v20 int32 = 1
            _ = v20
        } else {
            delete(v2, (v5 + v5))
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 uint16 = v3
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: "\n\t", F1: (v3 - (v3 * (^(v3)))), F2: "\n\t"}
                var v25 uint8 = 251
                _ = fi21
                _ = v24
                _ = v25
            }
            v4 = v4
        }
    }
    v2[v5] = "foo bar"
    v26 := v2[(v5 >> ((681) & 63))]
    var v27 bool = (0 >= 1)
    var v28 uint16 = 32767
    var v29 uint32 = 475
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

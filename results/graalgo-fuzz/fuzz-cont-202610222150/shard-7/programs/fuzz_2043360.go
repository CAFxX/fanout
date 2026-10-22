package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int64; F2 int }

func main() {
    var v1 int = -370
    v2 := []int{v1, 32767, (v1 ^ (^(v1)))}
    var v3 string = (("foo bar" + "foo bar") + "hello")
    var v4 uint64 = 6
    var v5 uint16 = 1
    fmt.Println(v1, v3)
    var v6 uint = 65535
    if (false && (3.14159 <= 0.1)) {
        v6 = v6
        var v7 string = v3
        _ = v7
    } else {
        _ = len(v2)
    }
    var v8 float64 = -2.25
    v4 = 42
    if ((v1 << (((v6 / (((v6 | v6)) | 1))) & 63)) == (v1 + int(v4))) {
        v9 := GenS1{F0: v1}
        _ = v9
    } else {
        switch ((v3 + v3) + "a") {
        case "foo bar":
            v10 := GenS2{F0: 927, F1: v3}
            _ = v10
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v5)
                for fi12 := 0; fi12 < 2; fi12++ {
                    if ((v4 <= (v4 << ((815) & 63))) || (466 <= (^(v4)))) {
                        if false {
                            v8 = ((v8 / (float64(-2.25) + v8)) - v8)
                            if false {
                                if true {
                                    v13 := GenS2{F0: 274, F1: "hello"}
                                    v14 := GenS2{F0: 4294967295, F1: v3}
                                    fmt.Println(v5, v4)
                                    _ = v13
                                    _ = v14
                                }
                            } else {
                                v3 = v3
                                fmt.Println(v3, v8, v3)
                            }
                        }
                    }
                    _ = fi12
                }
                v3 = ""
                _ = fi11
            }
        }
        if true {
            v4 = (v4 | v4)
            v15 := GenS3{F0: 42}
            switch "" {
            case "":
                var v16 uint = v6
                _ = v16
            case "hello":
                var v17 int8 = 127
                _ = v17
            case "\n\t":
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS4{F0: v3}
                    switch fi18 {
                    case -9223372036854775808:
                        if (1 > 1) {
                            var v20 int64 = 42
                            _ = v20
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = v15
        } else {
            v21 := GenS5{F0: (v1 + v1), F1: -557, F2: v1}
            var v22 float64 = (v8 * 3.14159)
            var v23 int32 = 42
            _ = v21
            _ = v22
            _ = v23
        }
        v2 = append(v2, (-(v1)))
    }
    v24 := GenS2{F0: 0, F1: v3}
    v4 += 18446744073709551615
    switch (v3 + (v3 + "foo bar")) {
    case "hello":
        var v25 int32 = 1
        _ = v25
    }
    fmt.Println(v6, v3, v6)
    _ = len(v2)
    var v26 uint = v6
    v27 := GenS2{F0: 1, F1: ((v3 + "a") + "\n\t")}
    var v28 uint16 = uint16(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []uint32{376, 878, 1, 0}
    var v2 int8 = -61
    var v3 int32 = 42
    v4 := []string{"foo bar", "\n\t", "foo bar"}
    var v5 int32 = -383
    v4 = append(v4, "\n\t")
    v4 = append(v4, ("a" + ""))
    for fi6 := 0; fi6 < 8; fi6++ {
        v2 += 42
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
    }
    var v7 int8 = (v2 & v2)
    for fi8 := 0; fi8 < 16; fi8++ {
        v2 = (v2 / (((v2 ^ v2)) | 1))
        v2 = (v7 << ((18446744073709551615) & 63))
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint = 42
            if (!true) {
                switch "foo bar" {
                case "hello":
                    v10 = v10
                }
                fmt.Println(v7, v3)
            } else {
                fmt.Println(v7, v7, v2)
            }
            v7 = v2
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    switch "\n\t" {
    case "a":
        for fi11 := 0; fi11 < 3; fi11++ {
            if ((-2.25 * float32(1e10)) <= (3.14159 + float32(-2.25))) {
                for fi12 := 0; fi12 < 10; fi12++ {
                    v7 = v7
                    v13 := GenS1{F0: 9223372036854775807}
                    _ = fi12
                    _ = v13
                }
            } else {
                var v14 int16 = -695
                var v15 int32 = (-(v3))
                v5 += v15
                _ = v14
                _ = v15
            }
            if true {
                v16 := GenS1{F0: 959}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch ("" + ("a" + "foo bar")) {
                case "\n\t":
                    if true {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v7 = v2
                            fmt.Println(fi11, v7, v3)
                            var v18 int32 = (v3 * (v3 - (v3 & (-(v5)))))
                            _ = fi17
                            _ = v18
                        }
                    }
                case "":
                    v2 = -61
                }
                _ = v16
            } else {
                var v19 int64 = 1
                _ = v19
            }
            _ = fi11
        }
    }
    var v20 uint64 = 18446744073709551615
    v5 -= (-(v3))
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint64 = (v20 + v20)
        var v23 bool = true
        if v23 {
            for fi24 := 0; fi24 < 19; fi24++ {
                v22 = v22
                v23 = v23
                _ = fi24
            }
            v5 = v5
            if ((v22 * v22) >= 18446744073709551615) {
                fmt.Println(v20, fi21)
                for fi25 := 0; fi25 < 20; fi25++ {
                    v23 = (((((fi21 ^ (fi25 ^ fi21)) != (fi21 % (((fi25 - 128)) | 1))) && v23) || true) && (42 == 2))
                    _ = fi25
                }
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS2{F0: v5}
        _ = fi26
        _ = v27
    }
    v5 = v3
    var v28 float64 = 1e10
    v29 := GenS2{F0: v3}
    fmt.Println(v28)
    var v30 bool = (false || false)
    v3 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 bool = (true || true)
    v2 := make(map[string]uint32)
    var v3 uint8 = 213
    var v4 bool = (!v1)
    v5 := GenS1{F0: 9223372036854775807, F1: "a", F2: 864}
    fmt.Println(v1, v1, v3)
    delete(v2, "\n\t")
    v6 := GenS2{F0: v3, F1: -32768}
    delete(v2, "\n\t")
    if (639 != 1) {
        for fi7 := 0; fi7 < 5; fi7++ {
            if (-128 > 17) {
                switch (fi7 ^ (fi7 + 1)) {
                case 42:
                    fmt.Println(v1, v1, v1)
                case -506:
                    switch fi7 {
                    case 128:
                        if false {
                            if (v3 >= 128) {
                                v8 := GenS3{F0: 344}
                                var v9 float64 = (3.14159 / 1.5)
                                var v10 int8 = 100
                                _ = v10
                                _ = v8
                                _ = v9
                            }
                            fmt.Println(fi7, v4, v3)
                            if true {
                                for fi11 := 0; fi11 < 2; fi11++ {
                                    v1 = v1
                                    _ = fi11
                                }
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    v4 = (423 <= 0)
                                case "a":
                                    var v12 string = "\n\t"
                                    _ = v12
                                default:
                                    for fi13 := 0; fi13 < 18; fi13++ {
                                        if (v4 || true) {
                                            v14 := GenS1{F0: -9223372036854775808, F1: "\n\t", F2: (^(fi13))}
                                            _ = v14
                                        }
                                        v15 := []int{}
                                        switch fi13 {
                                        case 1:
                                            v4 = v4
                                        case -9223372036854775808:
                                            var v16 float32 = 3.14159
                                            _ = v16
                                        case 410:
                                            v4 = true
                                        }
                                        _ = fi13
                                        _ = v15
                                    }
                                }
                            } else {
                                var v17 uint64 = 42
                                v4 = (true && (v4 && (100 >= 0)))
                                switch (("foo bar" + "foo bar") + ("a" + "hello")) {
                                case "\n\t":
                                    v18 := v2["hello"]
                                    _ = v18
                                case "":
                                    delete(v2, ("\n\t" + (("" + "") + "a")))
                                case "foo bar":
                                    var v19 uint = 255
                                    _ = v19
                                }
                                _ = v17
                            }
                        }
                    default:
                        if (true || (float32(3.14159) == (float32(1e10) * 1e10))) {
                            v20 := GenS1{F0: -266, F1: "\n\t", F2: 2147483647}
                            v21 := GenS3{F0: 65535}
                            _ = v20
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 7; fi22++ {
                                v23 := GenS2{F0: v3, F1: 42}
                                v3 = v3
                                _ = fi22
                                _ = v23
                            }
                        }
                    }
                case -9223372036854775808:
                    var v24 bool = (true || (!(1e10 <= 1.5)))
                    _ = v24
                }
            } else {
                var v25 int16 = -32768
                delete(v2, ("hello" + "a"))
                _ = v25
            }
            var v26 uint32 = 216
            fmt.Println(v4, v26, fi7)
            _ = fi7
            _ = v26
        }
    } else {
        v4 = v4
        var v27 string = "\n\t"
        for fi28 := 0; fi28 < 19; fi28++ {
            v4 = v1
            if (v3 < v3) {
                v29 := GenS1{F0: -9223372036854775808, F1: v27, F2: (^(fi28))}
                _ = v29
            }
            v1 = (!v4)
            _ = fi28
        }
        _ = v27
    }
    var v30 uint64 = 18446744073709551615
    var v31 uint64 = 692
    var v32 bool = (0 == 126)
    v33 := GenS4{F0: 100}
    var v34 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}

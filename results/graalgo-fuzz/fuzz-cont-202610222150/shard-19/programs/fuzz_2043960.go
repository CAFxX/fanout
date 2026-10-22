package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        switch "hello" {
        case "a":
            v8 := []int{}
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint8 = 110
                _ = fi9
                _ = v10
            }
        case "":
            var v11 uint8 = 255
            _ = v11
        default:
            if true {
                v12 := GenS1{F0: fi7}
                var v13 string = "a"
                v13 = "hello"
                _ = v12
                _ = v13
            } else {
                var v14 int = fi7
                v15 := []int{}
                var v16 bool = (-654 != 0)
                _ = v14
                _ = v15
                _ = v16
            }
        }
        _ = fi7
    }
    var v17 int64 = -9223372036854775808
    var v18 int16 = 1
    if ((!((!true) && (266 >= 1))) && (0.1 == (float32(0.1) * 0.0))) {
        for fi19 := 0; fi19 < 4; fi19++ {
            if true {
                v17 += (v17 | v17)
                v20 := GenS1{F0: 0}
                _ = v20
            } else {
                var v21 int8 = -1
                _ = v21
            }
            _ = fi19
        }
    }
    _ = v17
    _ = v18
    return -149
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v4 uint64 = 13
    var v5 uint = 408
    v6 := []string{("a" + ""), ("a" + ("\n\t" + (("\n\t" + "hello") + "a"))), (("a" + "a") + ""), "hello"}
    v22 := GenS1{F0: 128}
    if true {
        for fi23 := 0; fi23 < 6; fi23++ {
            if (!(4294967295 == 330)) {
                v2[("" + "a")] = 42
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 uint = v5
                    v4 = (v4 >> ((18446744073709551615) & 63))
                    if false {
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                        v25 = 2
                    } else {
                        if (!((v25 >> ((v25) & 63)) == 18446744073709551615)) {
                            var v26 int8 = 65
                            var v27 bool = (v26 >= (v26 - 99))
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(v5, v5)
                            v5 -= (v5 * v25)
                        }
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                fmt.Println(fi23, fi23, fi23)
                if true {
                    v28 := GenS2{F0: 1, F1: "\n\t", F2: (v5 % ((v5) | 1))}
                    v5 += 18446744073709551615
                    v4 = (v4 << ((v5) & 63))
                    _ = v28
                }
                v4 += (v4 & v4)
            }
            if true {
                if (true && (0 == -1)) {
                    v4 -= (v4 | v4)
                    v29 := GenS1{F0: fi23}
                    var v30 int64 = 0
                    _ = v29
                    _ = v30
                }
                var v31 uint32 = 4294967295
                var v32 bool = true
                _ = v31
                _ = v32
            } else {
                if (2 != 757) {
                    var v33 string = "a"
                    _ = v33
                }
                if true {
                    if true {
                        v34 := GenS3{F0: 0}
                        v4 -= v4
                        v35 := GenS4{F0: -1}
                        _ = v34
                        _ = v35
                    } else {
                        delete(v3, "a")
                    }
                } else {
                    v6 = append(v6, "\n\t")
                }
                var v36 int16 = 0
                _ = v36
            }
            _ = fi23
        }
        var v37 bool = true
        _ = v37
    } else {
        if false {
            delete(v1, 910)
            v4 += (v4 & 2)
            fmt.Println(v5, v5)
        }
    }
    v38 := GenS1{F0: 222}
    var v39 string = "hello"
    if false {
        _ = len(v1)
    }
    v40 := GenS3{F0: 219}
    if (0.1 != -2.25) {
        var v41 uint64 = (v4 >> ((v5) & 63))
        v4 = v4
        _ = v41
    }
    var v42 int8 = 116
    v43 := GenS3{F0: -9223372036854775808}
    var v44 float32 = 1.5
    _ = len(v2)
    var v45 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}

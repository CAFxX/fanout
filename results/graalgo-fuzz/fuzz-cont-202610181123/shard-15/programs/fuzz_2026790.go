package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint64, p1 uint64, p2 int8) uint16 {
    var v7 int = 975
    var v8 uint8 = 2
    v9 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() int64 {
    var v10 int32 = 0
    fmt.Println(v10)
    v10 += v10
    _ = v10
    return 579
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint32 = v1
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    var v5 int8 = 0
    var v6 uint16 = 65535
    var v11 int = 452
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint64 = 702
        var v14 float32 = 3.14159
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 = (v1 | (v1 / ((v2) | 1)))
    v15 := GenS1{F0: true}
    v16 := []int{}
    if ("\n\t" < "foo bar") {
        var v17 int16 = 1000
        if len(v16) > 0 {
            v16[0] = 606
        }
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v2 += v1
        _ = fi18
    }
    var v19 float64 = 0.0
    v16 = append(v16, 2147483647)
    var v20 int = v11
    for fi21 := 0; fi21 < 17; fi21++ {
        if (0.0 <= 3.14159) {
            if len(v16) > 0 {
                v16[0] = (v20 / ((598) | 1))
            }
        } else {
            var v22 int64 = int64(v1)
            fmt.Println(v1, v5)
            _ = v22
        }
        if (4294967295 == 42) {
            v23 := GenS2{F0: (v1 % ((42) | 1)), F1: (^(v6)), F2: 298}
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 string = "\n\t"
                var v26 uint64 = 1
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        } else {
            if (v19 > v19) {
                fmt.Println(v19, v20)
                v16 = append(v16, int(v6))
                v16 = append(v16, int(v6))
            } else {
                v5 = v5
                v19 += v19
                switch ("\n\t" + "hello") {
                case "foo bar":
                    v5 -= 127
                case "\n\t":
                    var v27 float32 = 1.5
                    _ = v27
                default:
                    v28 := GenS2{F0: v2, F1: (v6 * (^(v6))), F2: 0}
                    _ = v28
                }
            }
            switch "foo bar" {
            case "":
                for fi29 := 0; fi29 < 16; fi29++ {
                    if true {
                        if false {
                            switch ("foo bar" + ("" + ("a" + ((("foo bar" + "foo bar") + "foo bar") + ("\n\t" + ""))))) {
                            case "a":
                                fmt.Println(fi21)
                                var v30 int64 = genFn1()
                                v5 = (v5 % ((v5) | 1))
                                var v31 int = (fi29 ^ -2147483648)
                                v32 := GenS2{F0: v2, F1: v6, F2: 0}
                                _ = v30
                                _ = v31
                                _ = v32
                            case "hello":
                                fmt.Println(v20, fi29, v2)
                            case "foo bar":
                                var v33 int16 = 32767
                                _ = v33
                            }
                            switch (fi29 | 32768) {
                            case 128:
                                _ = len(v16)
                            case 32768:
                                fmt.Println(v6, v5)
                            }
                            if (false || (false && (v2 != v2))) {
                                v16 = append(v16, (v11 % (((-(v20))) | 1)))
                                v34 := GenS3{F0: "hello", F1: -546}
                                var v35 uint32 = v2
                                v4["a"] = v35
                                _ = v34
                                _ = v35
                            }
                        }
                        var v36 bool = false
                        v20 -= v11
                        v20 += 42
                        var v37 int = fi29
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int16 = -551
                        v6 += (^(v6))
                        _ = v38
                    }
                    _ = fi29
                }
            }
            var v39 int8 = v5
            _ = v39
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            if true {
                if (4294967295 < v2) {
                    v1 -= v2
                    _ = len(v3)
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                } else {
                    v16 = append(v16, (v20 >> ((4294967295) & 63)))
                    var v41 bool = true
                    _ = v41
                }
            }
            _ = fi40
        }
        _ = fi21
    }
    var v42 uint32 = 0
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}

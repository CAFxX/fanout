package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int64; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 int32) string {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint16 = 32767
        v8 := GenS1{F0: v7}
        var v9 bool = false
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            p0 = 1
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := []int{(fi10 - fi10), 32768, (-(fi10)), 0}
                if false {
                    if false {
                        if false {
                            if false {
                                if true {
                                    p0 += (^(p0))
                                }
                                v13 := GenS2{F0: -1}
                                _ = v13
                            }
                            if true {
                                v14 := GenS1{F0: 1}
                                _ = v14
                            } else {
                                if false {
                                    v15 := GenS3{F0: 255}
                                    var v16 uint8 = 128
                                    _ = v15
                                    _ = v16
                                } else {
                                    if (0 > 65535) {
                                        v17 := GenS1{F0: 2}
                                        _ = v17
                                    }
                                    v12 = append(v12, (fi10 % ((fi10) | 1)))
                                    v18 := GenS4{F0: 1}
                                    _ = v18
                                }
                            }
                            p0 -= (p0 ^ 623)
                        } else {
                            if (42 > 766) {
                                if len(v12) > 0 {
                                    v12[0] = fi10
                                }
                                p0 = (-(p0))
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            } else {
                                if true {
                                    v19 := GenS5{F0: 0, F1: 1, F2: 0}
                                    fmt.Println(p0)
                                    _ = v19
                                }
                            }
                        }
                        fmt.Println(fi10, fi10)
                        v20 := GenS6{F0: 1000}
                        _ = v20
                    } else {
                        var v21 uint32 = 4294967295
                        var v22 uint16 = 65535
                        if (42 < 42) {
                            if ((!true) || false) {
                                p0 -= p0
                                v23 := GenS7{F0: 697, F1: (^(fi10)), F2: 90}
                                fmt.Println(v22, p0, v22)
                                _ = v23
                            }
                        }
                        _ = v21
                        _ = v22
                    }
                    p0 = p0
                    if ((fi10 / ((fi10) | 1)) <= fi10) {
                        v24 := GenS4{F0: 42}
                        switch "" {
                        case "":
                            fmt.Println(fi10)
                        }
                        var v25 float32 = float32((3.14159 + ((-2.25 - -2.25) * 3.14159)))
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS1{F0: 255}
                        v27 := GenS3{F0: 490}
                        _ = v26
                        _ = v27
                    }
                } else {
                    v28 := GenS5{F0: 4294967295, F1: 9223372036854775807, F2: 1}
                    _ = v28
                }
                v12 = append(v12, fi10)
                _ = fi11
                _ = v12
            }
            var v29 int = (-(fi10))
            _ = v29
        } else {
            v30 := GenS1{F0: 32767}
            if true {
                switch (fi10 | fi10) {
                case 1:
                    for fi31 := 0; fi31 < 15; fi31++ {
                        p0 = (p0 & p0)
                        p0 += p0
                        p0 = (-(p0))
                        v32 := GenS7{F0: 691, F1: 42, F2: -1}
                        _ = fi31
                        _ = v32
                    }
                }
                var v33 uint8 = 128
                v33 = 255
                _ = v33
            } else {
                p0 += (p0 | -1)
                p0 += p0
            }
            _ = v30
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := []int{}
            v36 := GenS7{F0: 1, F1: 42, F2: 42}
            _ = len(v35)
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi10
    }
    _ = p0
    return "\n\t"
}
func genFn1(p0 string, p1 int16) string {
    p0 = p0
    fmt.Println(p0, p0)
    var v37 int32 = -601
    _ = p0
    _ = p1
    _ = v37
    return p0
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 uint16 = 145
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    v5 := []int{42, 42, 32767}
    fmt.Println(v2)
    v38 := GenS1{F0: (v2 ^ 116)}
    _ = len(v5)
    for fi39 := 0; fi39 < 14; fi39++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v40 := v3[0]
        var v41 int8 = 42
        _ = fi39
        _ = v40
        _ = v41
    }
    v2 -= (v2 & 2)
    v42 := GenS1{F0: (v2 * v2)}
    _ = len(v5)
    v2 = v2
    v2 -= (^(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}

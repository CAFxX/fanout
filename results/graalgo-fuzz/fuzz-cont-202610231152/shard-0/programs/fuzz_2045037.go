package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 uint = 42
    var v4 int16 = -32768
    v5 := []string{"a", "\n\t"}
    for fi6 := 0; fi6 < 11; fi6++ {
        _ = len(v2)
        v4 = 1000
        if true {
            v4 = v4
        } else {
            var v7 uint = v3
            if true {
                v8 := GenS1{F0: 0}
                fmt.Println(v3, v3, v3)
                _ = v8
            }
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v3, v4)
    v3 = (v3 + v3)
    if true {
        v3 = 677
        v9 := GenS2{F0: -2147483648, F1: 42}
        if (-9223372036854775808 >= -827) {
            if false {
                if true {
                    if true {
                        switch 1 {
                        case -9223372036854775808:
                            v3 += 42
                        case 42:
                            v4 -= v4
                        }
                        var v10 int32 = 983
                        _ = v10
                    }
                    v4 += v4
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v3 = v3
                        var v12 int32 = 2147483647
                        _ = fi11
                        _ = v12
                    }
                } else {
                    var v13 uint64 = 2
                    _ = v13
                }
                v4 = v4
            } else {
                v1[-848] = (1 != v4)
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 int32 = 42
                    _ = fi14
                    _ = v15
                }
            }
            if false {
                fmt.Println(v4, v3)
                v4 += v4
            } else {
                if false {
                    v3 -= v3
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    if false {
                        if true {
                            var v16 uint16 = 0
                            fmt.Println(v4)
                            _ = v16
                        }
                        v3 += v3
                    }
                    var v17 uint32 = 871
                    _ = v17
                }
            }
            v2[("\n\t" + ("hello" + "hello"))] = 0
        } else {
            v18 := GenS3{F0: 42, F1: 42}
            fmt.Println(v3, v3)
            _ = v18
        }
        _ = v9
    }
    var v19 float32 = float32(1e10)
    var v20 int32 = -610
    delete(v2, "foo bar")
    for fi21 := 0; fi21 < 18; fi21++ {
        _ = len(v5)
        if (2 == 213) {
            v22 := GenS3{F0: 42, F1: 299}
            fmt.Println(v3, fi21, v20)
            _ = v22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS2{F0: (v20 >> (((v3 & (^(v3)))) & 63)), F1: 220}
            fmt.Println(v3)
            var v25 uint64 = 42
            if true {
                v5 = append(v5, "")
                if false {
                    v4 -= v4
                }
                var v26 uint8 = 128
                _ = v26
            } else {
                if false {
                    v19 += v19
                    var v27 uint = (^(v3))
                    v19 = ((v19 * float32((0.1 - 1e10))) * (0.0 + (v19 - (0.0 + float32((3.14159 - 0.0))))))
                    _ = v27
                } else {
                    var v28 int32 = -1
                    fmt.Println(v19, fi23)
                    _ = v28
                }
            }
            var v29 uint32 = 767
            _ = fi23
            _ = v24
            _ = v25
            _ = v29
        }
        _ = fi21
    }
    var v30 int16 = (v4 % ((-1) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}

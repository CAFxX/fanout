package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint32) uint64 {
    var v7 int8 = 0
    v7 += int8(p1)
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 uint8) string {
    p0 = (^(p0))
    switch -796 {
    case 4294967295:
        v8 := []int{32768, 128, -1}
        _ = v8
    case 519:
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: -967}
            v11 := []int{(fi9 + -1), fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v12 := []int{-9223372036854775808, 9223372036854775807}
    p0 += 42
    _ = p0
    _ = p1
    _ = v12
    return "\n\t"
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 0
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := make(map[int]bool)
    v13 := GenS1{F0: -9223372036854775808}
    v5[9223372036854775807] = 4294967295
    switch "" {
    case "\n\t":
        _ = len(v4)
    case "hello":
        if false {
            if (1 < 872) {
                var v14 string = ("a" + (("\n\t" + "hello") + ("hello" + "a")))
                for fi15 := 0; fi15 < 15; fi15++ {
                    fmt.Println(v14, fi15)
                    switch ("\n\t" + v14) {
                    case "\n\t":
                        v14 = v14
                    }
                    switch (v14 + (v14 + v14)) {
                    case "":
                        fmt.Println(v14, v2)
                    case "hello":
                        v14 = "hello"
                    }
                    _ = fi15
                }
                if (2 <= 2) {
                    if (0 == 255) {
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v14 = (v14 + ("\n\t" + v14))
                            var v17 uint32 = uint32(v2)
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 int32 = 2147483647
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        var v20 uint8 = 25
                        switch 813 {
                        case 712:
                            var v21 int32 = 2147483647
                            _ = v21
                        case 1:
                            v22 := GenS1{F0: int64(v2)}
                            _ = v22
                        }
                        var v23 uint64 = 390
                        _ = v20
                        _ = v23
                    }
                } else {
                    v2 += (v2 + v2)
                }
                _ = v14
            }
            fmt.Println(v2)
        } else {
            var v24 float32 = (1.5 + 0.0)
            if true {
                v24 = 1e10
            }
            _ = v24
        }
    }
    var v25 uint = (v2 & v2)
    v26 := GenS1{F0: -1}
    var v27 int = 0
    if (3.14159 == 3.14159) {
        v27 = (v27 / ((42) | 1))
    }
    var v28 uint = (v25 >> ((1) & 63))
    fmt.Println(v2)
    var v29 int8 = -1
    var v30 int64 = -301
    var v31 uint32 = 2147483648
    if true {
        v32 := []int{(v27 + 2147483647)}
        if (42 < 1) {
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 uint64 = 229
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 uint8 = 0
                _ = fi35
                _ = v36
            }
        } else {
            if (42 > 923) {
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 bool = false
                    var v39 uint = v25
                    fmt.Println(v31)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                var v40 int64 = v30
                _ = v40
            } else {
                var v41 int = v27
                fmt.Println(v27, v31, v28)
                _ = v41
            }
        }
        v6[v27] = (255 == (^(v2)))
        _ = v32
    }
    var v42 uint = 1
    v43 := GenS1{F0: int64(v31)}
    fmt.Println(v25, v28, v30)
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}

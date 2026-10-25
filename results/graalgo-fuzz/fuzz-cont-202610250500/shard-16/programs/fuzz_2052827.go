package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 2
    var v3 bool = true
    v4 := make(map[uint32]string)
    var v5 int16 = 1000
    var v6 uint16 = 2
    if v3 {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: "foo bar", F1: (v2 >> ((4294967295) & 63)), F2: (v3 && (-1 > 1))}
            fmt.Println(v5, v1, fi7)
            var v9 float64 = 1.5
            _ = fi7
            _ = v8
            _ = v9
        }
        v5 += -777
    }
    if v3 {
        var v10 uint8 = 0
        var v11 uint16 = 32767
        _ = v10
        _ = v11
    } else {
        if ((v2 >> ((0) & 63)) < 293) {
            var v12 uint32 = 0
            if ((v1 * (v1 & 70)) <= v1) {
                v12 = (v12 * v12)
                if ((2 <= v12) && false) {
                    v13 := v4[1]
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := v4[v12]
                        _ = fi14
                        _ = v15
                    }
                    if (!v3) {
                        var v16 uint64 = (^(v2))
                        v3 = false
                        v17 := GenS1{F0: ("a" + ("hello" + ("\n\t" + ("foo bar" + "hello")))), F1: (v16 & v2), F2: (v16 < v2)}
                        _ = v16
                        _ = v17
                    } else {
                        v1 = (v1 ^ (v1 * v1))
                        v12 = v12
                        var v18 uint32 = (v12 & v12)
                        _ = v18
                    }
                    if v3 {
                        v19 := []int{-550, 0, -300, 4294967295}
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = v19
                    }
                }
                v12 = (v12 / ((2147483647) | 1))
            }
            if v3 {
                v20 := []int{-1, -825, 1, -9223372036854775808}
                _ = v20
            } else {
                var v21 int64 = -9223372036854775808
                if (0.1 < (0.0 / (-2.25 / (3.14159 / 1e10)))) {
                    var v22 uint = uint(v2)
                    switch 9223372036854775807 {
                    case 1:
                        for fi23 := 0; fi23 < 12; fi23++ {
                            for fi24 := 0; fi24 < 9; fi24++ {
                                fmt.Println(v5, v1)
                                v25 := GenS1{F0: "hello", F1: (^(v2)), F2: ("hello" < "a")}
                                fmt.Println(v3, v5)
                                _ = fi24
                                _ = v25
                            }
                            v5 -= v5
                            _ = fi23
                        }
                    }
                    if (!(!v3)) {
                        v26 := []int{183, -456, int(v21)}
                        _ = v26
                    } else {
                        fmt.Println(v2, v1)
                    }
                    _ = v22
                }
                _ = v21
            }
            _ = v12
        } else {
            v2 -= (v2 - (v2 / ((0) | 1)))
            fmt.Println(v6, v2)
            var v27 int8 = v1
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v1 = (-(v1))
        if (1 == 65535) {
            delete(v4, 83)
            delete(v4, 2147483648)
            v29 := []int{1, fi28}
            _ = v29
        } else {
            var v30 uint64 = 643
            _ = v30
        }
        _ = fi28
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        fmt.Println(v5, v2)
        _ = fi31
    }
    fmt.Println(v2)
    fmt.Println(v1)
    var v32 float64 = (0.1 + (1.5 - 3.14159))
    for fi33 := 0; fi33 < 16; fi33++ {
        v1 = 42
        _ = fi33
    }
    v4[2] = "a"
    v34 := GenS1{F0: "hello", F1: uint64(v6), F2: false}
    var v35 bool = false
    v3 = (682 != 412)
    v35 = false
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 7; fi37++ {
            switch -1 {
            case -378:
                v38 := []int{(^(fi36))}
                _ = v38
            case -713:
                v39 := []int{fi36}
                _ = v39
            }
            v32 = v32
            _ = fi37
        }
        var v40 uint64 = uint64(v6)
        _ = fi36
        _ = v40
    }
    switch 9223372036854775807 {
    case 804:
        v5 += 42
    case 42:
        var v41 string = "a"
        _ = v41
    case 1:
        v32 = v32
    }
    v42 := GenS2{F0: 0, F1: v2}
    v43 := []int{-392, -1}
    var v44 uint32 = uint32(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}

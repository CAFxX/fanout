package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16; F1 uint; F2 int32 }
func genFn0(p0 uint, p1 uint8) uint16 {
    var v6 int8 = 121
    v7 := GenS1{F0: -32768, F1: 4294967295}
    v6 = v6
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 bool) int {
    if false {
        p0 = ((p0 && p0) && (!true))
        p0 = (78 < 42)
    } else {
        fmt.Println(p0)
    }
    v9 := GenS1{F0: -32768, F1: 995}
    p0 = p0
    if false {
        var v10 uint16 = 149
        if true {
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(p0, v10, p0)
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            fmt.Println(v10, v10)
            var v13 uint16 = 37
            _ = v13
        }
        _ = v10
    }
    var v14 uint8 = 42
    _ = p0
    _ = v14
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    var v3 int64 = -9223372036854775808
    var v4 uint16 = uint16(v1)
    var v5 int8 = -128
    switch genFn1(false) {
    case 127:
        v15 := GenS1{F0: 1, F1: 2147483648}
        _ = v15
    case -833:
        switch 32768 {
        case 59:
            var v16 uint16 = genFn0(4294967295, 42)
            _ = v16
        case -2147483648:
            if false {
                if false {
                    var v17 float64 = 3.14159
                    v3 -= (v3 * 73)
                    var v18 int8 = (v5 ^ (^(v5)))
                    _ = v17
                    _ = v18
                }
            }
        case 0:
            v19 := GenS1{F0: 1000, F1: uint32(v1)}
            _ = v19
        }
    case 42:
        switch 1 {
        case -9223372036854775808:
            delete(v2, "a")
        case 80:
            var v20 uint = 384
            _ = v20
        }
    }
    v21 := []int{-752, 409, 127, 1}
    v3 += (^(v3))
    if (!(926 > 32768)) {
        if true {
            switch 9223372036854775807 {
            case -190:
                if ("a" > "\n\t") {
                    v22 := GenS1{F0: 296, F1: 2147483647}
                    for fi23 := 0; fi23 < 5; fi23++ {
                        v1 += 128
                        _ = fi23
                    }
                    if true {
                        v3 = -308
                        if (true && true) {
                            delete(v2, ("\n\t" + "a"))
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v25 := GenS2{F0: "foo bar", F1: (fi24 << ((175) & 63))}
                                _ = len(v21)
                                _ = fi24
                                _ = v25
                            }
                        } else {
                            _ = len(v21)
                        }
                    }
                    _ = v22
                }
            case 127:
                var v26 bool = (1e10 >= 0.1)
                _ = v26
            }
        }
        var v27 uint = 4294967295
        var v28 uint64 = 2
        _ = v27
        _ = v28
    }
    var v29 int8 = -113
    var v30 float32 = 1e10
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS2{F0: "\n\t", F1: (fi31 | genFn1(false))}
        for fi33 := 0; fi33 < 3; fi33++ {
            switch (-(fi31)) {
            case -649:
                var v34 bool = true
                _ = v34
            }
            var v35 int32 = 1
            _ = fi33
            _ = v35
        }
        fmt.Println(v29)
        _ = fi31
        _ = v32
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v2[""] = 118
        _ = fi36
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v38 := GenS2{F0: "", F1: (fi37 ^ genFn1(true))}
        v39 := GenS3{F0: -998, F1: 152, F2: -2147483648}
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 string = "\n\t"
    v29 = (v29 >> ((915) & 63))
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 uint8 = v1
        if false {
            var v43 int8 = (v5 % ((v5) | 1))
            v4 = v4
            var v44 uint16 = 0
            _ = v43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    fmt.Println(v30)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v5
}

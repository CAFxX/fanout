package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int32, p1 uint32, p2 uint) uint64 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint = p2
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p1, p0)
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 uint64) uint8 {
    v8 := GenS1{F0: 0, F1: -49}
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(p0, p1)
        var v10 int8 = 42
        v11 := GenS1{F0: 65535, F1: v10}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 65535, F1: 100}
        var v14 int64 = -442
        var v15 int32 = 2147483647
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    var v3 bool = (false && false)
    v4 := make(map[uint32]bool)
    if v3 {
        var v16 uint8 = 1
        if v3 {
            v17 := GenS1{F0: 65535, F1: -128}
            v18 := GenS1{F0: 2, F1: 127}
            _ = v17
            _ = v18
        } else {
            if false {
                var v19 int64 = -9223372036854775808
                _ = v19
            } else {
                if len(v2) > 0 {
                    v2[0] = ("foo bar" + "")
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 string = ""
                    var v22 int16 = 1000
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            if v3 {
                if true {
                    var v23 uint64 = 451
                    if v3 {
                        var v24 uint8 = 0
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 16; fi25++ {
                            fmt.Println(v3, v16, v23)
                            v16 -= genFn1(((v16 > 2) && v3), v23)
                            _ = fi25
                        }
                    }
                    _ = v23
                }
            } else {
                switch (("\n\t" + "a") + "hello") {
                case "foo bar":
                    v16 -= v16
                }
                var v26 int = 483
                var v27 int16 = 1000
                _ = v26
                _ = v27
            }
        }
        _ = v16
    }
    var v28 int32 = -1
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if false {
        var v29 bool = v3
        _ = v29
    } else {
        delete(v1, 0)
        v3 = true
    }
    if v3 {
        v30 := GenS2{F0: v3}
        if (!v3) {
            v31 := GenS2{F0: (905 == -1)}
            v3 = ("hello" < "\n\t")
            _ = v31
        } else {
            for fi32 := 0; fi32 < 14; fi32++ {
                v3 = v3
                v3 = (-2.25 <= 3.14159)
                v3 = false
                _ = fi32
            }
            _ = len(v2)
            var v33 int = -2147483648
            _ = v33
        }
        var v34 uint64 = genFn0(2147483647, 1, 277)
        _ = v30
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi35
    }
    var v36 uint32 = 420
    var v37 int64 = 51
    for fi38 := 0; fi38 < 14; fi38++ {
        var v39 uint = 1
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}

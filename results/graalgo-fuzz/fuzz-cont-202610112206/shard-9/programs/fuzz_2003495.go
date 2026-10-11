package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    var v1 int64 = 0
    var v2 int32 = -1
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 int16 = -1
    switch "" {
    case "a":
        var v6 int32 = (v2 << ((65535) & 63))
        _ = v6
    case "":
        v7 := GenS1{F0: 2147483647, F1: 42}
        _ = v7
    case "hello":
        var v8 uint = 4294967295
        _ = v8
    default:
        v2 += (^(v2))
    }
    v9 := GenS1{F0: -990, F1: 0}
    var v10 uint32 = 65535
    fmt.Println(v10, v2, v5)
    delete(v4, ("a" + ("" + ("" + ("hello" + "hello")))))
    switch 2147483647 {
    case 0:
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint8 = 128
                if true {
                    fmt.Println(v5, v1)
                    v14 := []int{}
                    var v15 uint64 = 0
                    _ = v14
                    _ = v15
                }
                var v16 float64 = (1e10 + 0.1)
                var v17 float32 = (1e10 / 3.14159)
                _ = fi12
                _ = v13
                _ = v16
                _ = v17
            }
            _ = fi11
        }
    }
    if false {
        v2 = v2
        if true {
            var v18 int = -996
            _ = v18
        } else {
            v1 += (v1 % ((42) | 1))
        }
        v19 := GenS2{F0: v2}
        _ = v19
    }
    if true {
        if (42 > 255) {
            switch ("foo bar" + "") {
            case "foo bar":
                if (true && false) {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v5 = 42
                        v5 += v5
                        _ = fi20
                    }
                    var v21 int64 = (^(v1))
                    _ = v21
                } else {
                    var v22 int = 167
                    switch ("a" + "\n\t") {
                    case "hello":
                        var v23 uint = 0
                        _ = v23
                    case "a":
                        switch (v22 & (v22 | 349)) {
                        case 0:
                            var v24 string = ("" + "a")
                            _ = v24
                        case 32767:
                            _ = len(v4)
                        case 994:
                            v2 = -657
                        }
                    }
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 uint16 = v3
                        var v27 int = fi25
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = v22
                }
            case "a":
                var v28 int64 = v1
                _ = v28
            case "hello":
                v29 := GenS1{F0: 32768, F1: 65535}
                _ = v29
            default:
                v30 := []int{}
                _ = v30
            }
            var v31 int8 = 42
            _ = v31
        } else {
            for fi32 := 0; fi32 < 16; fi32++ {
                v33 := GenS3{F0: 65535, F1: (fi32 / (((^(fi32))) | 1))}
                switch "" {
                case "":
                    var v34 int = fi32
                    _ = v34
                case "foo bar":
                    for fi35 := 0; fi35 < 19; fi35++ {
                        if true {
                            var v36 int = fi32
                            v37 := GenS2{F0: (v2 ^ -1)}
                            _ = v36
                            _ = v37
                        } else {
                            v10 += v10
                            v38 := GenS2{F0: 0}
                            _ = v38
                        }
                        switch (fi35 + (-(fi32))) {
                        case -9223372036854775808:
                            v39 := []int{(fi35 >> ((2) & 63)), 42, fi32, 4294967295}
                            _ = v39
                        default:
                            var v40 bool = ((v10 / (((v10 + (v10 % ((10) | 1)))) | 1)) == 65535)
                            _ = v40
                        }
                        if false {
                            v5 = (v5 - (^(v5)))
                        } else {
                            fmt.Println(v5, v5)
                        }
                        _ = fi35
                    }
                }
                var v41 float32 = 1e10
                _ = fi32
                _ = v33
                _ = v41
            }
            v3 = (v3 >> ((932) & 63))
        }
        if true {
            v42 := []int{}
            _ = v42
        } else {
            var v43 string = ("" + "hello")
            if true {
                switch 32768 {
                case -63:
                    v3 = (v3 << ((42) & 63))
                }
            }
            v2 -= -1
            _ = v43
        }
        for fi44 := 0; fi44 < 6; fi44++ {
            v1 = (v1 - (v1 / (((v1 << ((255) & 63))) | 1)))
            v45 := []int{1, (fi44 + fi44), fi44}
            v46 := GenS3{F0: v3, F1: (fi44 & -1)}
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    var v47 int8 = 42
    fmt.Println(v10, v2, v3)
    var v48 int = int(v10)
    var v49 uint = 454
    fmt.Println(v10, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v9
}

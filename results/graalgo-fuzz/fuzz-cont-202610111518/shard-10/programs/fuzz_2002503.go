package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int8 }
func genFn0(p0 uint8) uint32 {
    if false {
        var v6 uint64 = 4294967295
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = p0
        switch ((("\n\t" + "hello") + "\n\t") + ("foo bar" + "a")) {
        case "":
            v8 := []int{}
            _ = v8
        case "hello":
            if true {
                v9 := GenS1{F0: 42}
                if (((("" + "a") + "a") + "foo bar") < "foo bar") {
                    p0 = (p0 ^ (p0 * 1))
                } else {
                    if (-44 >= 337) {
                        v10 := GenS2{F0: 42, F1: 32767, F2: 166}
                        v11 := GenS1{F0: -128}
                        _ = v10
                        _ = v11
                    }
                }
                _ = v9
            } else {
                v12 := []int{-260}
                v13 := GenS2{F0: 942, F1: 0, F2: 42}
                _ = v12
                _ = v13
            }
        case "a":
            v14 := []int{9223372036854775807}
            _ = v14
        }
        _ = fi7
    }
    _ = p0
    return 65535
}
func genFn1(p0 uint) uint32 {
    v15 := []int{128, -1}
    v15 = append(v15, 4294967295)
    _ = p0
    _ = v15
    return genFn0(2)
}

func main() {
    var v1 uint64 = 807
    v2 := []string{}
    var v3 bool = true
    var v4 string = ""
    var v5 string = v4
    v16 := GenS2{F0: 42, F1: 775, F2: 65535}
    v17 := GenS2{F0: 0, F1: 613, F2: 384}
    v1 = v1
    var v18 int32 = -728
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 uint16 = 2
    switch (v5 + ("\n\t" + "foo bar")) {
    case "a":
        var v20 int16 = -32768
        _ = v20
    case "\n\t":
        var v21 int16 = -525
        _ = v21
    case "":
        switch "a" {
        case "foo bar":
            v22 := GenS2{F0: (v19 | (v19 >> ((42) & 63))), F1: 727, F2: (v19 % ((v19) | 1))}
            _ = v22
        case "a":
            var v23 bool = (42 <= 22)
            _ = v23
        default:
            if false {
                fmt.Println(v1)
                if (v19 > v19) {
                    v24 := GenS1{F0: 127}
                    var v25 int64 = -1
                    _ = v24
                    _ = v25
                } else {
                    v3 = v3
                    var v26 int64 = 42
                    v19 = 255
                    _ = v26
                }
                v5 = v4
            }
        }
    }
    if (!(1 <= 672)) {
        v27 := GenS3{F0: 42}
        var v28 int32 = 42
        _ = v27
        _ = v28
    }
    var v29 uint = 18446744073709551615
    if (!(117 >= 100)) {
        v18 -= 0
        if v3 {
            if false {
                _ = len(v2)
                v5 = "a"
                v30 := GenS4{F0: 666}
                _ = v30
            } else {
                for fi31 := 0; fi31 < 9; fi31++ {
                    fmt.Println(v1, v3, v1)
                    v32 := GenS1{F0: -111}
                    _ = fi31
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int32 = (v18 - v18)
                v29 = v29
                _ = fi33
                _ = v34
            }
        } else {
            fmt.Println(v18)
            _ = len(v2)
        }
        fmt.Println(v5, v3)
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        fmt.Println(v4, v19)
        fmt.Println(fi35, v3, v19)
        _ = fi35
    }
    v18 = (v18 + -987)
    for fi36 := 0; fi36 < 4; fi36++ {
        for fi37 := 0; fi37 < 18; fi37++ {
            fmt.Println(v1, v18)
            switch -9223372036854775808 {
            case 766:
                var v38 int16 = -32768
                _ = v38
            case 32767:
                v39 := GenS5{F0: 1, F1: -41}
                _ = v39
            }
            if v3 {
                v29 = uint(v1)
            } else {
                var v40 uint8 = 255
                _ = v40
            }
            _ = fi37
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}

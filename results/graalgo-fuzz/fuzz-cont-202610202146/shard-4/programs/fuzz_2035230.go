package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int8; F2 uint16 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 string = "foo bar"
        _ = fi6
        _ = v7
    }
    var v8 bool = (-128 != -1)
    v9 := GenS1{F0: 0, F1: -1, F2: 42}
    v8 = v8
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int32 = -2147483648
        var v12 string = ("hello" + ("hello" + "hello"))
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v8
    _ = v9
    return 1000
}
func genFn1(p0 int16) int8 {
    if (1 <= int(p0)) {
        switch 128 {
        case -9223372036854775808:
            p0 -= 1
        case 2147483647:
            if false {
                fmt.Println(p0)
            }
        case 32767:
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int = fi13
                var v15 int32 = -1
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        v16 := GenS2{F0: ""}
        p0 -= p0
        _ = v16
    } else {
        v17 := []int{9223372036854775807}
        _ = v17
    }
    if false {
        p0 = (p0 | p0)
        var v18 int32 = 42
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: 1, F2: -1}
    _ = p0
    _ = v19
    return 1
}

func main() {
    var v1 int32 = 809
    var v2 int8 = -94
    var v3 int = 7
    v4 := []uint32{825}
    v5 := make(map[int]int)
    switch ("foo bar" + "") {
    case "hello":
        var v20 string = ""
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        fmt.Println(v1)
        fmt.Println(v1)
        _ = fi21
    }
    v22 := GenS3{F0: 0}
    var v23 uint32 = 2147483647
    if true {
        v2 = v2
    }
    v24 := GenS1{F0: int64(v2), F1: 1, F2: genFn1(-844)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v25 uint16 = 255
    switch ("a" + "hello") {
    case "hello":
        if ((int64(v3) != -1) || ((v23 << ((42) & 63)) > (v23 * 4294967295))) {
            if (((3.14159 * 0.1) - (((((1e10 / 1e10) + -2.25) / -2.25) / 1.5) + (3.14159 * (float64(float32((1e10 + 1e10))) + (0.1 + -2.25))))) > float64(0.1)) {
                v26 := GenS3{F0: int64(v25)}
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 int16 = 32767
                    var v29 uint16 = 0
                    v30 := v5[fi27]
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = v26
            }
            v1 -= (v1 & 1)
            if false {
                var v31 bool = ((!(844 >= 0)) && ("a" != "\n\t"))
                _ = v31
            }
        }
    case "a":
        var v32 bool = (1 > 1)
        _ = v32
    case "\n\t":
        var v33 uint32 = 766
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        _ = len(v4)
        v35 := GenS4{F0: 1, F1: (v2 >> ((4294967295) & 63)), F2: (^(v25))}
        _ = fi34
        _ = v35
    }
    if (uint8(v25) < 0) {
        for fi36 := 0; fi36 < 8; fi36++ {
            fmt.Println(v2, fi36, v3)
            fmt.Println(v23, v3, fi36)
            _ = fi36
        }
    }
    v3 = v3
    v3 = (v3 | (v3 << ((65535) & 63)))
    fmt.Println(v23)
    v2 = genFn1(42)
    var v37 uint = 2
    for fi38 := 0; fi38 < 16; fi38++ {
        for fi39 := 0; fi39 < 1; fi39++ {
            switch (fi39 % ((fi38) | 1)) {
            case 32768:
                v2 = 1
            case 9223372036854775807:
                var v40 int16 = -1
                _ = v40
            case 428:
                var v41 uint16 = v25
                _ = v41
            }
            _ = fi39
        }
        v2 += v2
        var v42 string = "foo bar"
        _ = fi38
        _ = v42
    }
    v43 := GenS3{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v5
}

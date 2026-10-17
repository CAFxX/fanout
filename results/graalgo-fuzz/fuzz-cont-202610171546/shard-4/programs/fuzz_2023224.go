package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int32; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 uint; F2 int8 }
func genFn0(p0 uint32) int32 {
    p0 = p0
    v5 := []int{1, -9223372036854775808, 4294967295}
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            var v7 int8 = 100
            for fi8 := 0; fi8 < 19; fi8++ {
                v7 = v7
                _ = fi8
            }
            _ = v7
        } else {
            var v9 int16 = 1000
            var v10 bool = (("hello" != "hello") || true)
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: -128}
        v13 := GenS1{F0: 127}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 uint32) int8 {
    p1 = (^(p1))
    switch "" {
    case "\n\t":
        p1 -= (p1 << ((824) & 63))
    case "":
        var v14 uint64 = 0
        _ = v14
    case "foo bar":
        p1 = 1
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 629}
    var v18 int32 = (p0 | 42)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := []string{"\n\t"}
    v4 := []uint32{694, 65535, 4294967295}
    if true {
        var v19 uint32 = 2147483648
        v20 := GenS3{F0: -596, F1: -1, F2: "a"}
        _ = v19
        _ = v20
    } else {
        v21 := GenS4{F0: 4294967295}
        if (203 != 0) {
            var v22 string = ((("\n\t" + "") + "\n\t") + "\n\t")
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v22
        }
        var v23 uint32 = 65535
        _ = v21
        _ = v23
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v24 bool = (!false)
        _ = v24
    case "a":
        for fi25 := 0; fi25 < 9; fi25++ {
            switch fi25 {
            case -946:
                var v26 string = "a"
                _ = v26
            case 32767:
                if true {
                    var v27 int32 = -307
                    v27 = genFn0(547)
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    _ = v27
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi25
        }
    case "":
        var v28 int16 = 1000
        _ = v28
    }
    if false {
        v29 := v1[0]
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 float64 = 1e10
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := v2["hello"]
            var v34 int32 = 42
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v29
    }
    var v35 uint32 = 42
    switch -155 {
    case 1:
        if (!true) {
            switch "foo bar" {
            case "a":
                var v36 uint8 = 0
                _ = v36
            }
        }
    }
    if false {
        v37 := GenS5{F0: 4294967295}
        v38 := GenS6{F0: (2 <= 2)}
        _ = v37
        _ = v38
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v35 = v35
        for fi40 := 0; fi40 < 14; fi40++ {
            if false {
                v3 = append(v3, ("a" + "hello"))
                fmt.Println(fi39, fi39, v35)
                v35 += 2147483647
            } else {
                fmt.Println(v35, fi39)
                if true {
                    v41 := GenS7{F0: 2, F1: 42, F2: genFn1(-1, (v35 * 637))}
                    _ = v41
                }
            }
            fmt.Println(fi39, fi39)
            v35 += v35
            _ = fi40
        }
        _ = fi39
    }
    v35 = (v35 << ((255) & 63))
    var v42 int16 = -668
    for fi43 := 0; fi43 < 16; fi43++ {
        v44 := GenS1{F0: -128}
        var v45 float32 = 1.5
        _ = fi43
        _ = v44
        _ = v45
    }
    v46 := GenS2{F0: 65535}
    var v47 int64 = -809
    v48 := GenS2{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v48
}

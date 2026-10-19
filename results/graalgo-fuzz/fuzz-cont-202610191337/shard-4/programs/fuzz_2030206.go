package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
func genFn0() int16 {
    var v7 uint64 = 42
    if false {
        v8 := GenS1{F0: 42, F1: -250}
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS2{F0: 18446744073709551615}
            switch fi9 {
            case 0:
                fmt.Println(fi9, fi9)
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        var v11 int8 = 100
        if (244 == 4294967295) {
            v12 := []int{170, 1}
            _ = v12
        }
        _ = v11
    }
    if true {
        fmt.Println(v7)
        var v13 int16 = -257
        _ = v13
    } else {
        var v14 int16 = -1
        v7 += (v7 - v7)
        v15 := []int{42}
        _ = v14
        _ = v15
    }
    var v16 int8 = 100
    _ = v16
    _ = v7
    return 32767
}
func genFn1(p0 int, p1 bool) uint {
    switch (p0 >> ((829) & 63)) {
    case 9223372036854775807:
        fmt.Println(p0, p1)
    case -1:
        for fi17 := 0; fi17 < 12; fi17++ {
            if (1e10 > 3.14159) {
                var v18 int64 = -9223372036854775808
                var v19 uint = 255
                var v20 int32 = -2147483648
                _ = v18
                _ = v19
                _ = v20
            }
            if (-1 > int32(p0)) {
                v21 := []int{}
                var v22 int16 = 42
                _ = v21
                _ = v22
            } else {
                var v23 uint64 = 279
                v23 = 1
                v24 := GenS2{F0: 42}
                _ = v23
                _ = v24
            }
            var v25 float64 = -2.25
            _ = fi17
            _ = v25
        }
    case 42:
        switch (((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + ("\n\t" + "hello")) + "a") {
        case "foo bar":
            v26 := []int{-572}
            _ = v26
        case "hello":
            var v27 int64 = -719
            _ = v27
        }
    }
    if (p1 || p1) {
        v28 := []int{44, 0}
        p0 = p0
        _ = v28
    } else {
        if false {
            v29 := GenS1{F0: 506, F1: -200}
            fmt.Println(p1, p1)
            _ = v29
        }
        switch (p0 << ((65535) & 63)) {
        case 0:
            var v30 int = p0
            _ = v30
        case 32768:
            if ((4294967295 < 401) || p1) {
                switch "a" {
                case "":
                    v31 := GenS3{F0: -32768, F1: 595, F2: -62}
                    _ = v31
                case "foo bar":
                    p0 = -32
                }
                var v32 uint16 = 42
                _ = v32
            }
        }
        p0 -= 834
    }
    switch ("foo bar" + "a") {
    case "a":
        var v33 uint64 = 0
        _ = v33
    case "\n\t":
        if (18446744073709551615 <= 42) {
            var v34 uint = 65535
            v35 := GenS2{F0: (v34 & v34)}
            _ = v34
            _ = v35
        } else {
            var v36 int64 = 0
            var v37 int32 = -997
            _ = v36
            _ = v37
        }
    case "":
        if (p1 || (2147483647 < 2147483647)) {
            p0 = -151
        }
    default:
        if (32767 < -784) {
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint16 = 65535
                p0 = (^(fi38))
                _ = fi38
                _ = v39
            }
            v40 := GenS4{F0: -911, F1: 65535, F2: 2}
            v41 := []int{(p0 + (^(p0)))}
            _ = v40
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        v43 := GenS3{F0: -393, F1: 42, F2: 76}
        _ = fi42
        _ = v43
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = ("\n\t" + "hello")
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    v6 := []uint32{630, 1, 108}
    var v44 uint32 = 0
    _ = len(v6)
    var v45 uint32 = v44
    switch "" {
    case "foo bar":
        v3[v1] = (v44 - (^(v44)))
    case "":
        fmt.Println(v44)
    case "a":
        var v46 int32 = 866
        _ = v46
    }
    var v47 string = (v1 + "\n\t")
    var v48 int8 = -1
    switch v47 {
    case "":
        v1 = ((v1 + ("\n\t" + (v47 + "hello"))) + "")
    case "a":
        switch v47 {
        case "":
            fmt.Println(v44, v44)
        }
    }
    var v49 int = -2147483648
    delete(v5, (v1 + v47))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}

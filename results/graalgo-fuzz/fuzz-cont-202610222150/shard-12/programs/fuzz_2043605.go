package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 uint; F2 string }
func genFn0(p0 uint16) int8 {
    switch "a" {
    case "":
        if (true || (1 < 18446744073709551615)) {
            v5 := GenS1{F0: -9223372036854775808}
            _ = v5
        } else {
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 uint32 = 65535
                v8 := []int{32768, fi6}
                p0 = p0
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 float32 = 3.14159
            _ = v9
        }
    case "hello":
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint = 4294967295
            v12 := GenS1{F0: -882}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 32767}
            v16 := []int{int(p0), fi14, fi14, 2147483647}
            _ = len(v16)
            v17 := GenS1{F0: 1}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        switch -634 {
        case 42:
            var v18 uint64 = 4294967295
            _ = v18
        case 1:
            p0 = (p0 * 122)
        }
        var v19 int = (fi13 | -1)
        _ = fi13
        _ = v19
    }
    fmt.Println(p0)
    fmt.Println(p0)
    var v20 uint = 2
    _ = p0
    _ = v20
    return -1
}

func main() {
    var v1 bool = true
    var v2 int = 352
    var v3 uint64 = 42
    var v4 int = (v2 | v2)
    switch "" {
    case "a":
        for fi21 := 0; fi21 < 9; fi21++ {
            fmt.Println(v4, fi21, v3)
            _ = fi21
        }
    case "hello":
        v1 = v1
    }
    v1 = ((v3 & 2) != v3)
    if v1 {
        var v22 float32 = 1e10
        v23 := GenS3{F0: v3, F1: 265, F2: -128}
        v2 = v4
        _ = v22
        _ = v23
    } else {
        v3 = (v3 >> ((255) & 63))
        v24 := []int{9223372036854775807, v4}
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        if (1 != 205) {
            switch ("\n\t" + "hello") {
            case "":
                v3 = v3
            }
        }
        var v26 int8 = 0
        v27 := GenS2{F0: 903}
        _ = fi25
        _ = v26
        _ = v27
    }
    v2 = (v2 * v4)
    switch "" {
    case "\n\t":
        switch 128 {
        case 32767:
            v3 = 0
        }
    case "foo bar":
        if true {
            v28 := GenS4{F0: "hello"}
            _ = v28
        }
    }
    var v29 int16 = 1000
    for fi30 := 0; fi30 < 11; fi30++ {
        switch v4 {
        case 0:
            v3 = v3
        case 1:
            fmt.Println(v4)
        case 42:
            v29 = v29
        default:
            switch "" {
            case "a":
                v2 += v2
            case "hello":
                v2 -= v4
            case "\n\t":
                v31 := []int{(v4 % ((9223372036854775807) | 1)), 177}
                _ = v31
            default:
                if false {
                    var v32 int8 = 127
                    var v33 int = (v2 | v4)
                    _ = v32
                    _ = v33
                }
            }
        }
        if false {
            v34 := GenS5{F0: v29}
            v35 := []int{fi30}
            if v1 {
                switch ("hello" + ("foo bar" + "hello")) {
                case "\n\t":
                    fmt.Println(v29, v1, v4)
                }
                var v36 uint8 = 239
                var v37 uint32 = 2
                _ = v36
                _ = v37
            } else {
                v38 := GenS6{F0: v3, F1: 18446744073709551615, F2: ("\n\t" + "\n\t")}
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 int16 = (v29 << ((4294967295) & 63))
                    v29 = v29
                    var v41 float32 = 1e10
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = v38
            }
            _ = v34
            _ = v35
        } else {
            for fi42 := 0; fi42 < 15; fi42++ {
                fmt.Println(v4, fi42, v2)
                var v43 int16 = (v29 >> ((961) & 63))
                v3 -= (v3 / (((v3 * v3)) | 1))
                v44 := []int{}
                v2 -= fi30
                _ = fi42
                _ = v43
                _ = v44
            }
            v2 += (v4 - (fi30 ^ (fi30 ^ 32767)))
        }
        _ = fi30
    }
    v45 := GenS5{F0: 1}
    for fi46 := 0; fi46 < 19; fi46++ {
        if (true && v1) {
            v47 := []int{}
            var v48 uint16 = 556
            _ = v47
            _ = v48
        } else {
            var v49 uint = 0
            _ = v49
        }
        _ = fi46
    }
    v50 := []int{}
    v2 = (^(v2))
    if (v1 && false) {
        var v51 uint64 = v3
        _ = v51
    } else {
        if false {
            v1 = (0.0 >= (1.5 - -2.25))
            if len(v50) >= 2 {
                v50 = v50[1:len(v50)]
            }
        } else {
            var v52 string = "a"
            var v53 int16 = v29
            v4 -= v2
            _ = v52
            _ = v53
        }
        if len(v50) > 0 {
            v50[0] = v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v50
}

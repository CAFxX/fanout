package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int32; F2 int }
func genFn0() int16 {
    var v4 int32 = -801
    switch -1 {
    case 4294967295:
        var v5 uint8 = 238
        _ = v5
    default:
        if true {
            var v6 bool = (0 == -105)
            _ = v6
        }
    }
    var v7 int32 = v4
    v7 = v7
    v4 = 1
    _ = v4
    _ = v7
    return 32767
}
func genFn1(p0 uint8, p1 int16, p2 uint32) uint {
    var v8 uint32 = p2
    if (p1 == p1) {
        if false {
            var v9 bool = false
            for fi10 := 0; fi10 < 8; fi10++ {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := GenS1{F0: fi11}
                    p1 = genFn0()
                    _ = fi11
                    _ = v12
                }
                v9 = true
                if ((v9 && v9) && ((true && true) || ((0 != (p0 << ((42) & 63))) && (3.14159 >= 1e10)))) {
                    var v13 int16 = genFn0()
                    var v14 int16 = -1
                    _ = v13
                    _ = v14
                } else {
                    var v15 uint8 = p0
                    v16 := GenS2{F0: "a"}
                    _ = v15
                    _ = v16
                }
                _ = fi10
            }
            switch ("hello" + "foo bar") {
            case "foo bar":
                fmt.Println(v9, v8, p1)
            case "\n\t":
                switch "a" {
                case "\n\t":
                    fmt.Println(p2, p0)
                }
            }
            _ = v9
        }
        v17 := []int{-990}
        var v18 uint64 = 0
        _ = v17
        _ = v18
    } else {
        fmt.Println(p2)
    }
    v19 := []int{1}
    var v20 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]string)
    var v3 int8 = -73
    v21 := GenS2{F0: "hello"}
    v22 := []int{42}
    if true {
        if false {
            for fi23 := 0; fi23 < 11; fi23++ {
                v3 -= 127
                v1 += 492
                _ = fi23
            }
            var v24 uint64 = 18446744073709551615
            _ = v24
        } else {
            if true {
                v3 = v3
                fmt.Println(v1, v3)
            }
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            switch (("" + ("\n\t" + "a")) + "hello") {
            case "hello":
                v3 -= -124
            case "foo bar":
                fmt.Println(v1)
            }
            _ = fi25
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            switch (("" + "\n\t") + "a") {
            case "":
                var v27 int64 = 42
                _ = v27
            case "foo bar":
                for fi28 := 0; fi28 < 11; fi28++ {
                    if (0 <= 32767) {
                        var v29 int8 = v3
                        _ = v29
                    }
                    fmt.Println(v3, fi28, fi28)
                    v1 = 0
                    _ = fi28
                }
            case "a":
                for fi30 := 0; fi30 < 2; fi30++ {
                    fmt.Println(v1, fi30, fi30)
                    var v31 uint64 = 903
                    delete(v2, (^(fi30)))
                    _ = fi30
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 5; fi32++ {
                var v33 float32 = 1.5
                _ = fi32
                _ = v33
            }
            fmt.Println(fi26, v1)
            _ = fi26
        }
    }
    v34 := GenS1{F0: -1}
    for fi35 := 0; fi35 < 13; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            v2[(-(fi36))] = "hello"
            fmt.Println(fi36)
            v3 = v3
            _ = fi36
        }
        _ = fi35
    }
    v37 := GenS3{F0: 607}
    for fi38 := 0; fi38 < 4; fi38++ {
        v3 -= (v3 | 100)
        v3 = (v3 & (v3 << ((18446744073709551615) & 63)))
        _ = fi38
    }
    if ((2147483647 > -974) && (false && true)) {
        var v39 bool = true
        _ = v39
    } else {
        switch 42 {
        case 4294967295:
            v40 := GenS4{F0: 32768, F1: -2147483648, F2: -9223372036854775808}
            _ = v40
        case 640:
            if true {
                fmt.Println(v3, v3)
            } else {
                v41 := GenS3{F0: 18446744073709551615}
                _ = v41
            }
        default:
            delete(v2, 128)
        }
        switch 523 {
        case 42:
            for fi42 := 0; fi42 < 10; fi42++ {
                for fi43 := 0; fi43 < 10; fi43++ {
                    v2[128] = "a"
                    var v44 float32 = -2.25
                    var v45 float32 = 0.0
                    switch (fi43 + -1) {
                    case 42:
                        var v46 int8 = 100
                        _ = v46
                    default:
                        var v47 int64 = 999
                        _ = v47
                    }
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                switch (("a" + ("hello" + "foo bar")) + ("hello" + "\n\t")) {
                case "\n\t":
                    v48 := GenS1{F0: fi42}
                    _ = v48
                default:
                    for fi49 := 0; fi49 < 12; fi49++ {
                        var v50 uint = 0
                        _ = fi49
                        _ = v50
                    }
                }
                var v51 int = 42
                _ = fi42
                _ = v51
            }
        case 127:
            v52 := GenS2{F0: ("a" + ("foo bar" + "hello"))}
            _ = v52
        }
        for fi53 := 0; fi53 < 8; fi53++ {
            var v54 int64 = -9223372036854775808
            if false {
                for fi55 := 0; fi55 < 5; fi55++ {
                    v3 = v3
                    _ = fi55
                }
                fmt.Println(v3, fi53)
                v1 = 486
            } else {
                var v56 string = "foo bar"
                switch (fi53 % (((fi53 << ((2) & 63))) | 1)) {
                case -2147483648:
                    var v57 uint8 = 2
                    _ = v57
                case -84:
                    v58 := GenS3{F0: 680}
                    _ = v58
                default:
                    delete(v2, 2147483647)
                }
                _ = v56
            }
            if false {
                v54 = v54
            }
            _ = fi53
            _ = v54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v37
}

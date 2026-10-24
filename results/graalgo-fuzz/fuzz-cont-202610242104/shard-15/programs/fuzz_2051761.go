package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int64 }
func genFn0() string {
    switch ("foo bar" + "a") {
    case "foo bar":
        var v4 float64 = 1e10
        _ = v4
    case "\n\t":
        v5 := GenS1{F0: -334}
        _ = v5
    }
    v6 := []int{421, 127}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch 321 {
        case 128:
            v8 := GenS1{F0: 42}
            _ = v8
        case -180:
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v9 int8 = 127
                _ = v9
            } else {
                var v10 int32 = -2147483648
                if true {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := GenS1{F0: -1}
                        var v13 uint = 255
                        var v14 uint = (v13 << ((255) & 63))
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS2{F0: 18446744073709551615}
                    var v16 uint32 = 833
                    _ = v15
                    _ = v16
                }
                fmt.Println(v10)
                _ = v10
            }
        case -9223372036854775808:
            for fi17 := 0; fi17 < 4; fi17++ {
                switch "\n\t" {
                case "\n\t":
                    var v18 uint32 = 990
                    _ = v18
                }
                var v19 string = "\n\t"
                switch (fi7 << ((42) & 63)) {
                case -768:
                    fmt.Println(v19, fi7)
                }
                v19 = v19
                _ = fi17
                _ = v19
            }
        }
        _ = fi7
    }
    var v20 int32 = -387
    fmt.Println(v20)
    _ = v20
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 42
    var v3 int64 = 788
    switch -491 {
    case 32767:
        v21 := GenS3{F0: 611, F1: -778}
        _ = v21
    default:
        if true {
            v3 = (v3 * v3)
            v2 += (v2 >> ((406) & 63))
            delete(v1, 4294967295)
        } else {
            v22 := GenS2{F0: 2}
            switch 2147483647 {
            case 127:
                var v23 int64 = (v3 % (((v3 & (v3 * (v3 / ((v3) | 1))))) | 1))
                _ = v23
            default:
                if false {
                    fmt.Println(v2, v2)
                }
            }
            switch ("\n\t" + ("hello" + "hello")) {
            case "":
                v24 := []int{-770, -678, 0, 1}
                _ = v24
            }
            _ = v22
        }
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 uint16 = 372
            _ = fi26
            _ = v27
        }
        var v28 int64 = (v3 / (((-(v3))) | 1))
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 int = 1
            _ = fi29
            _ = v30
        }
        _ = fi25
        _ = v28
    }
    var v31 uint32 = 1
    v32 := GenS1{F0: 42}
    v33 := GenS4{F0: 128}
    for fi34 := 0; fi34 < 3; fi34++ {
        v3 = v3
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v31, v31, fi34)
        case "foo bar":
            v2 = (v2 + v2)
        }
        _ = fi34
    }
    switch (("foo bar" + "foo bar") + "a") {
    case "foo bar":
        v35 := GenS5{F0: 101, F1: v3, F2: 32767}
        _ = v35
    case "hello":
        v2 += (v2 ^ (v2 % ((42) | 1)))
    case "\n\t":
        v31 = (^(v31))
    default:
        v36 := GenS2{F0: 42}
        _ = v36
    }
    if false {
        var v37 int16 = (v2 ^ (v2 / (((v2 / ((v2) | 1))) | 1)))
        v31 += 1
        var v38 string = "a"
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 bool = true
            _ = fi39
            _ = v40
        }
        v41 := []int{-9223372036854775808}
        _ = v41
    }
    fmt.Println(v31, v31)
    fmt.Println(v31, v31, v31)
    for fi42 := 0; fi42 < 16; fi42++ {
        _ = len(v1)
        v31 = (^(v31))
        var v43 string = ("hello" + ("a" + ""))
        _ = fi42
        _ = v43
    }
    var v44 int64 = (v3 * v3)
    v31 = 2147483647
    if false {
        v45 := GenS5{F0: 121, F1: 0, F2: 255}
        _ = v45
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        v47 := []int{}
        _ = fi46
        _ = v47
    }
    v44 = v3
    v48 := GenS6{F0: 1}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v44
    _ = v48
}

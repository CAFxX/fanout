package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 uint) int64 {
    var v5 string = "\n\t"
    v5 = ((v5 + "\n\t") + v5)
    if (316 == (p1 / (((p1 / ((p1) | 1))) | 1))) {
        v6 := []int{}
        if (-782 == 9223372036854775807) {
            v5 = v5
            var v7 int64 = -464
            _ = v7
        } else {
            if p0 {
                var v8 uint16 = 32767
                _ = v8
            }
            var v9 uint = p1
            fmt.Println(v9)
            _ = v9
        }
        _ = v6
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if p0 {
            var v11 bool = (true && (2 < 8))
            _ = v11
        } else {
            fmt.Println(p0)
            var v12 uint8 = 127
            var v13 float64 = (1e10 * (0.1 / -2.25))
            _ = v12
            _ = v13
        }
        switch int(p1) {
        case -601:
            var v14 string = v5
            _ = v14
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v5
    return 558
}

func main() {
    var v1 uint8 = 76
    var v2 int16 = 1000
    var v3 uint64 = 42
    v4 := make(map[int]bool)
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int16 = 1
        _ = fi15
        _ = v16
    }
    delete(v4, 0)
    var v17 int = 32768
    var v18 uint16 = 65535
    v17 = v17
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int8 = 1
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v18)
    v3 += v3
    v21 := GenS1{F0: 42, F1: ("hello" <= "a")}
    var v22 uint8 = (^(v1))
    v2 = 117
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            v22 = v22
            v24 := []int{127}
            v1 = (v1 % ((v22) | 1))
            _ = v24
        }
        _ = fi23
    }
    if ((v18 * v18) == (v18 | v18)) {
        if true {
            switch "hello" {
            case "\n\t":
                switch ("\n\t" + "foo bar") {
                case "a":
                    switch (v17 >> ((4294967295) & 63)) {
                    case -1:
                        _ = len(v4)
                    default:
                        v25 := GenS2{F0: 546}
                        _ = v25
                    }
                case "\n\t":
                    v26 := GenS3{F0: ("a" + "")}
                    _ = v26
                case "foo bar":
                    var v27 int32 = -2147483648
                    _ = v27
                default:
                    if false {
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v18 -= (v18 & v18)
                            v29 := GenS1{F0: -9223372036854775808, F1: (v2 >= v2)}
                            switch (("" + "\n\t") + "") {
                            case "":
                                fmt.Println(v18)
                            }
                            _ = fi28
                            _ = v29
                        }
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            fmt.Println(v2, v18)
                        }
                    } else {
                        var v30 uint64 = v3
                        _ = v30
                    }
                }
            case "hello":
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := GenS4{F0: 0, F1: 42, F2: -1}
                    _ = fi31
                    _ = v32
                }
            case "foo bar":
                fmt.Println(v17)
            default:
                fmt.Println(v17)
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                fmt.Println(v22, fi33)
                _ = fi33
            }
        }
        fmt.Println(v18, v3, v3)
        var v34 string = ("\n\t" + "a")
        _ = v34
    } else {
        v35 := GenS4{F0: -9223372036854775808, F1: (v1 / ((0) | 1)), F2: 42}
        v3 = v3
        fmt.Println(v3, v22)
        _ = v35
    }
    var v36 bool = (false && (!(0 != -1)))
    var v37 string = ("foo bar" + "hello")
    v2 = v2
    for fi38 := 0; fi38 < 20; fi38++ {
        v4[(-(fi38))] = v36
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 float32 = -2.25
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println(v18, v37)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int32, p1 int) int8 {
    if ((p1 * (p1 | p1)) > -1) {
        p1 += -624
    } else {
        if true {
            var v6 int8 = 0
            v6 -= v6
            switch "foo bar" {
            case "":
                v7 := GenS1{F0: false}
                _ = v7
            case "a":
                fmt.Println(v6, p1)
            }
            _ = v6
        }
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "foo bar":
            var v8 uint16 = 0
            _ = v8
        case "":
            v9 := GenS1{F0: false}
            _ = v9
        case "hello":
            v10 := []int{}
            _ = v10
        }
        if (!(!(-2.25 < 0.1))) {
            if (!(true || (!false))) {
                v11 := GenS1{F0: true}
                p1 += (-(p1))
                p0 = p0
                _ = v11
            }
            var v12 float64 = 3.14159
            _ = v12
        } else {
            if (false && (true || true)) {
                v13 := []int{1000}
                p0 -= 0
                p0 = (p0 % (((^(p0))) | 1))
                _ = v13
            }
            var v14 uint = 255
            _ = v14
        }
    }
    fmt.Println(p0, p0)
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint64 = 1
        if ((1000 <= 1000) && (1 <= 408)) {
            for fi17 := 0; fi17 < 17; fi17++ {
                if ((100 != -111) || (false || ((42 > 128) && (!true)))) {
                    fmt.Println(v16, fi15, fi15)
                }
                v16 -= v16
                _ = fi17
            }
        } else {
            v18 := GenS1{F0: true}
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    switch "\n\t" {
    case "foo bar":
        var v19 uint16 = 147
        _ = v19
    case "":
        fmt.Println(p0)
    }
    p1 += p1
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 bool = true
    v4 := make(map[int]int)
    v5 := []int{32768, 325, -57, 864}
    _ = len(v5)
    v3 = (30 < 128)
    v20 := GenS1{F0: (128 == -532)}
    if (!(620 <= 42)) {
        _ = len(v1)
    } else {
        v3 = ("hello" > ("foo bar" + ""))
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v3)
        }
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if v3 {
            var v22 int8 = -49
            _ = v22
        }
        _ = fi21
    }
    if v3 {
        switch ("" + "foo bar") {
        case "":
            var v23 uint = 822
            _ = v23
        case "foo bar":
            var v24 int64 = 0
            _ = v24
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v25 uint16 = 65535
        _ = v25
    default:
        fmt.Println(v3)
    }
    fmt.Println(v3)
    v3 = true
    v3 = v3
    var v26 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

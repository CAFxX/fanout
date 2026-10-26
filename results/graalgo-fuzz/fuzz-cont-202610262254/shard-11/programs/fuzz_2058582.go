package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 int64) string {
    switch "a" {
    case "a":
        p0 += 0
    case "hello":
        switch ("" + "") {
        case "foo bar":
            v4 := []int{2147483647, 377, 377}
            _ = v4
        case "":
            var v5 float32 = 1e10
            _ = v5
        }
    }
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: 247}
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    return ("hello" + "hello")
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 uint16 = 718
    var v7 int32 = 0
    v8 := GenS1{F0: 170}
    var v9 int8 = 1
    v7 -= (v7 >> ((4294967295) & 63))
    var v10 int64 = v2
    if true {
        if true {
            v9 -= (^(v9))
        }
        if (248 > 110) {
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(v7)
                _ = fi11
            }
        }
        switch (("hello" + "foo bar") + ("hello" + "\n\t")) {
        case "\n\t":
            fmt.Println(v2, v3, v9)
        default:
            v1[""] = true
        }
    } else {
        if ("" < "hello") {
            if (-183 == 864) {
                for fi12 := 0; fi12 < 13; fi12++ {
                    _ = len(v1)
                    for fi13 := 0; fi13 < 10; fi13++ {
                        fmt.Println(v7, fi12)
                        var v14 uint = 914
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                v7 -= v7
            }
        } else {
            if (1 <= 128) {
                v15 := []int{-425, 32768}
                _ = v15
            }
            var v16 uint16 = v3
            v9 = int8(v2)
            _ = v16
        }
        var v17 float32 = (0.0 + float32(1.5))
        var v18 string = "foo bar"
        _ = v17
        _ = v18
    }
    v9 = (v9 - 0)
    var v19 float32 = -2.25
    if (!true) {
        for fi20 := 0; fi20 < 15; fi20++ {
            if ("foo bar" < ("\n\t" + ("a" + ("" + ("a" + ("\n\t" + "hello")))))) {
                v21 := []int{fi20, (fi20 + fi20), int(v9), fi20}
                for fi22 := 0; fi22 < 10; fi22++ {
                    v3 += v3
                    v23 := GenS2{F0: (fi20 ^ fi20), F1: v10, F2: true}
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    var v25 int16 = -48
                    switch 128 {
                    case -506:
                        var v26 int16 = v25
                        _ = v26
                    default:
                        v27 := GenS2{F0: (^(fi20)), F1: v10, F2: false}
                        _ = v27
                    }
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    _ = fi24
                    _ = v25
                }
                _ = v21
            } else {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := v1["hello"]
                    v2 -= int64(fi28)
                    var v30 uint = 1
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            switch "" {
            case "\n\t":
                v31 := []int{1, fi20, fi20, (^(fi20))}
                _ = v31
            default:
                switch -960 {
                case 4294967295:
                    var v32 uint = 65535
                    _ = v32
                }
            }
            var v33 float32 = (-2.25 - ((v19 / v19) / v19))
            _ = fi20
            _ = v33
        }
        if (!(!false)) {
            var v34 bool = ((!(!(65535 >= 0))) || false)
            _ = v34
        }
        var v35 int16 = 697
        _ = v35
    }
    v7 = v7
    v2 += -174
    var v36 float32 = 0.0
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 int16 = 42
        v39 := GenS3{F0: -1}
        v7 -= -2
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v7
    _ = v8
    _ = v9
}

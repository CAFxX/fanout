package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int, p1 uint32) int {
    var v4 int64 = 9223372036854775807
    p1 = p1
    p1 = p1
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 float64 = ((1e10 / -2.25) - -2.25)
    v3 = ((v3 * ((v3 / (v3 / v3)) - v3)) - float64(1e10))
    v3 = (v3 / v3)
    if ((128 == 255) && (!((!true) && false))) {
        if false {
            v3 += (v3 + (v3 / ((0.0 * 3.14159) - float64(0.1))))
            v3 = (0.0 + 1e10)
            var v6 int64 = -1
            _ = v6
        } else {
            switch "a" {
            case "":
                var v7 uint8 = 41
                _ = v7
            }
            v8 := GenS1{F0: 1}
            if true {
                var v9 float64 = v3
                _ = len(v1)
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 uint16 = 658
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                v12 := GenS2{F0: (("foo bar" + ("a" + "\n\t")) + (("hello" + ("\n\t" + (("\n\t" + "hello") + (("hello" + "\n\t") + "")))) + ("a" + ((("a" + (((("a" + ("a" + "")) + ("a" + "hello")) + "") + "hello")) + "hello") + "")))), F1: 32767}
                var v13 int16 = -1
                var v14 uint8 = 1
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v8
        }
        v15 := []int{-646, 127, genFn0(1, 2147483648)}
        _ = v15
    }
    var v16 int16 = 1000
    v16 = (-(v16))
    var v17 uint64 = 656
    _ = len(v1)
    if (true && (("\n\t" + "foo bar") > ("foo bar" + ""))) {
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 string = ""
                _ = fi19
                _ = v20
            }
            fmt.Println(v17, fi18)
            v17 += v17
            _ = fi18
        }
        v21 := GenS1{F0: 32768}
        _ = v21
    } else {
        if (v17 > v17) {
            v22 := []int{}
            _ = v22
        } else {
            v17 = v17
            v17 = 42
        }
        if (-1 > -888) {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 uint64 = v17
                v25 := GenS1{F0: (fi23 - fi23)}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        switch 127 {
        case 32768:
            fmt.Println(v16)
        }
    }
    v17 -= v17
    v26 := []int{2147483647, 540, -207}
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 float64 = (float64(-2.25) * float64(-2.25))
        _ = fi27
        _ = v28
    }
    fmt.Println(v17)
    _ = len(v26)
    var v29 int8 = 42
    var v30 float32 = 1e10
    if false {
        if true {
            v2[("hello" + "a")] = (true && (37 == 393))
            for fi31 := 0; fi31 < 18; fi31++ {
                var v32 uint16 = 2
                for fi33 := 0; fi33 < 6; fi33++ {
                    v34 := GenS1{F0: 926}
                    _ = fi33
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
        } else {
            v17 = 250
            var v35 bool = true
            if (v35 && v35) {
                var v36 uint = 0
                fmt.Println(v35, v17, v16)
                _ = v36
            }
            _ = v35
        }
        var v37 uint8 = 2
        _ = v37
    } else {
        fmt.Println(v17)
        var v38 int8 = v29
        if ((v17 / ((v17) | 1)) < (^(v17))) {
            v39 := GenS3{F0: 88, F1: 4294967295}
            _ = v39
        }
        _ = v38
    }
    v40 := GenS2{F0: ("\n\t" + "foo bar"), F1: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v40
}

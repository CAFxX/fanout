package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
func genFn0() int16 {
    var v5 uint = 4294967295
    for fi6 := 0; fi6 < 3; fi6++ {
        v5 += v5
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 string = "\n\t"
            v8 = v8
            _ = fi7
            _ = v8
        }
        v9 := []int{(fi6 & fi6)}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v5 += 2
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 bool = (127 == 94)
            if true {
                v13 := []int{}
                if v12 {
                    switch (^(fi10)) {
                    case -126:
                        var v14 int64 = 0
                        _ = v14
                    case -941:
                        var v15 int64 = 1
                        _ = v15
                    }
                    if (-801 >= -287) {
                        var v16 uint = 65535
                        _ = v16
                    }
                }
                var v17 int8 = 127
                _ = v13
                _ = v17
            } else {
                v18 := []int{}
                switch "a" {
                case "foo bar":
                    var v19 float64 = 1.5
                    _ = v19
                case "hello":
                    v20 := GenS1{F0: 65535}
                    _ = v20
                case "a":
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                default:
                    v21 := GenS1{F0: 403}
                    _ = v21
                }
                switch ("\n\t" + "\n\t") {
                case "a":
                    v5 = (v5 % (((v5 >> ((v5) & 63))) | 1))
                case "foo bar":
                    v12 = (-1 == -80)
                case "\n\t":
                    fmt.Println(v12, fi11)
                }
                _ = v18
            }
            if v12 {
                v22 := []int{}
                v23 := GenS2{F0: "\n\t"}
                if len(v22) > 0 {
                    v22[0] = 1
                }
                v5 = 4294967295
                var v24 string = "foo bar"
                var v25 string = "foo bar"
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            } else {
                v26 := GenS1{F0: 42}
                var v27 uint16 = 0
                fmt.Println(v27, v27, v12)
                _ = v26
                _ = v27
            }
            _ = fi11
            _ = v12
        }
        v28 := GenS2{F0: ((("foo bar" + "a") + "hello") + ("foo bar" + ("" + ("foo bar" + "hello"))))}
        _ = fi10
        _ = v28
    }
    _ = v5
    return -441
}

func main() {
    var v1 float64 = float64(0.0)
    var v2 int64 = 1
    var v3 int = 128
    v4 := make(map[string]int)
    v29 := GenS2{F0: ("" + "foo bar")}
    v3 += (v3 >> ((42) & 63))
    v2 -= v2
    if false {
        var v30 uint = 2
        var v31 uint16 = 2
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        if true {
            var v33 float32 = (3.14159 * -2.25)
            v34 := GenS3{F0: 226}
            _ = v33
            _ = v34
        }
        _ = fi32
    }
    if (int16(v2) >= -32768) {
        v35 := GenS3{F0: 759}
        var v36 int = (v3 & 149)
        _ = v35
        _ = v36
    }
    v37 := []int{v3, v3, (v3 / ((v3) | 1))}
    v2 = (v2 % ((v2) | 1))
    switch "hello" {
    case "\n\t":
        var v38 float32 = float32(float64(3.14159))
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
}

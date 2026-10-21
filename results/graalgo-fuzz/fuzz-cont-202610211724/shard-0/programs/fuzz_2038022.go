package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
func genFn0(p0 int64, p1 uint, p2 uint) int32 {
    switch -2147483648 {
    case -362:
        p0 = p0
    case 366:
        v6 := []int{}
        _ = v6
    }
    v7 := GenS1{F0: "a", F1: "", F2: uint32(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int32) int16 {
    if true {
        var v8 int = 4294967295
        v9 := []int{v8, v8, (v8 | 32768), 32767}
        var v10 int8 = 127
        _ = v10
        _ = v8
        _ = v9
    }
    p0 -= 982
    v11 := []int{427, -175, 354, -1}
    if (!(115 >= 48)) {
        v12 := GenS1{F0: ("" + "a"), F1: ((("a" + "foo bar") + ("" + "")) + "foo bar"), F2: 0}
        _ = v12
    } else {
        p1 = (-(p1))
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 bool = true
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 uint64 = 4294967295
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int16 = 520
        _ = fi13
        _ = v14
    }
    fmt.Println(v5, v2)
    var v15 float64 = 0.1
    for fi16 := 0; fi16 < 10; fi16++ {
        _ = len(v3)
        v2 = (42 == 897)
        _ = fi16
    }
    if (32768 < -9223372036854775808) {
        v17 := []int{-2147483648, 32768, -890, 0}
        _ = v17
    } else {
        var v18 string = (("a" + "") + "")
        _ = v18
    }
    if (1 <= 65535) {
        v2 = v2
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int32 = -764
            _ = fi19
            _ = v20
        }
    } else {
        switch "a" {
        case "\n\t":
            var v21 int16 = 1
            _ = v21
        case "":
            delete(v4, ("foo bar" + "a"))
        }
        v15 = 3.14159
    }
    v22 := []int{9223372036854775807, 32768}
    _ = len(v3)
    var v23 uint = 132
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS1{F0: "hello", F1: ("foo bar" + "a"), F2: uint32(v23)}
        var v26 int16 = int16(v23)
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        if v2 {
            v1 = (v1 ^ (v1 + (v1 | v1)))
            var v28 uint8 = v1
            _ = v28
        } else {
            switch (("hello" + ("" + (("" + "a") + ""))) + ("\n\t" + ((("hello" + "a") + ("hello" + ("foo bar" + "\n\t"))) + (("hello" + (("foo bar" + (("a" + "hello") + ("\n\t" + (("\n\t" + "foo bar") + ("" + (("foo bar" + "") + "foo bar")))))) + ("a" + (("" + "\n\t") + "a")))) + "foo bar")))) {
            case "":
                var v29 uint8 = v1
                _ = v29
            case "\n\t":
                var v30 uint32 = 462
                _ = v30
            case "foo bar":
                v4[""] = ""
            }
        }
        _ = fi27
    }
    fmt.Println(v1, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() uint {
    var v6 int32 = 915
    v7 := GenS1{F0: 4294967295, F1: (!false)}
    v6 = (v6 - v6)
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := []string{}
    v2 := []int{1, 2147483647, -250, 32767}
    v3 := []string{"\n\t"}
    v4 := []string{"hello", (("hello" + ("a" + "foo bar")) + ("a" + ("\n\t" + "foo bar"))), "foo bar", ("foo bar" + ((("a" + "") + ("foo bar" + "\n\t")) + ("a" + "")))}
    var v5 bool = (!(!(false && (!false))))
    if (1000 < -435) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            v5 = false
        } else {
            var v8 uint = 870
            switch (("\n\t" + "hello") + (("" + "hello") + (("foo bar" + "foo bar") + ""))) {
            case "a":
                var v9 string = "foo bar"
                _ = v9
            case "\n\t":
                var v10 uint8 = 42
                _ = v10
            case "foo bar":
                var v11 uint64 = 18446744073709551615
                _ = v11
            }
            fmt.Println(v8)
            _ = v8
        }
    } else {
        v5 = v5
    }
    _ = len(v1)
    var v12 uint16 = 1
    v5 = v5
    var v13 string = "\n\t"
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: 4294967295, F1: false}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int16 = 42
            var v18 uint32 = 187
            var v19 uint = 0
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v13, v12, v13)
    var v20 float32 = 1e10
    v13 = (v13 + (v13 + v13))
    var v21 int8 = 1
    v22 := GenS2{F0: "", F1: -716}
    for fi23 := 0; fi23 < 3; fi23++ {
        if v5 {
            var v24 uint = 18446744073709551615
            fmt.Println(fi23, v5)
            var v25 int64 = -9223372036854775808
            _ = v24
            _ = v25
        }
        var v26 uint16 = v12
        var v27 int32 = -1
        _ = fi23
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
